package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

const (
	defaultConfigFilename = "torram-relayer.yml"
	defaultDataDirname    = "data"
)

var (
	defaultAppDataDir  = btcutil.AppDataDir("torram-relayer", false)
	defaultConfigFile  = filepath.Join(defaultAppDataDir, defaultConfigFilename)
	defaultRPCKeyFile  = filepath.Join(defaultAppDataDir, "rpc.key")
	defaultRPCCertFile = filepath.Join(defaultAppDataDir, "rpc.cert")
)

// DataDir returns the path to the data directory based on the provided home path.
func DataDir(homePath string) string {
	return filepath.Join(homePath, defaultDataDirname)
}

// Config defines the top-level configuration for the Torram Relayer.
type Config struct {
	Common    CommonConfig    `mapstructure:"common"`
	BTC       BTCConfig       `mapstructure:"btc"`
	GRPC      GRPCConfig      `mapstructure:"grpc"`
	GRPCWeb   GRPCWebConfig   `mapstructure:"grpc-web"`
	Metrics   MetricsConfig   `mapstructure:"metrics"`
	Submitter SubmitterConfig `mapstructure:"submitter"`
	Reporter  ReporterConfig  `mapstructure:"reporter"`
}

// Validate checks the configuration for any invalid or missing values.
func (cfg *Config) Validate() error {
	if err := cfg.Common.Validate(); err != nil {
		return fmt.Errorf("invalid config in common: %w", err)
	}

	if err := cfg.BTC.Validate(); err != nil {
		return fmt.Errorf("invalid config in BTC: %w", err)
	}

	if err := cfg.GRPC.Validate(); err != nil {
		return fmt.Errorf("invalid config in GRPC: %w", err)
	}

	if err := cfg.GRPCWeb.Validate(); err != nil {
		return fmt.Errorf("invalid config in GRPC-Web: %w", err)
	}

	if err := cfg.Metrics.Validate(); err != nil {
		return fmt.Errorf("invalid config in metrics: %w", err)
	}

	if err := cfg.Submitter.Validate(); err != nil {
		return fmt.Errorf("invalid config in submitter: %w", err)
	}

	if err := cfg.Reporter.Validate(); err != nil {
		return fmt.Errorf("invalid config in reporter: %w", err)
	}

	return nil
}

// CreateLogger initializes and returns a logger using the common configuration.
func (cfg *Config) CreateLogger() (*zap.Logger, error) {
	return cfg.Common.CreateLogger()
}

// DefaultConfigFile returns the default configuration file path.
func DefaultConfigFile() string {
	return defaultConfigFile
}

// DefaultConfig returns the default configuration for the Torram Relayer.
func DefaultConfig() *Config {
	return &Config{
		Common:    DefaultCommonConfig(),
		BTC:       DefaultBTCConfig(),
		GRPC:      DefaultGRPCConfig(),
		GRPCWeb:   DefaultGRPCWebConfig(),
		Metrics:   DefaultMetricsConfig(),
		Submitter: DefaultSubmitterConfig(),
		Reporter:  DefaultReporterConfig(),
	}
}

// New loads and parses the configuration from the specified file.
func New(configFile string) (Config, error) {
	if _, err := os.Stat(configFile); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// The specified config file does not exist.
			return Config{}, fmt.Errorf("no config file found at %s", configFile)
		}
		return Config{}, err // Other file access errors.
	}

	// Parse the configuration file.
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
