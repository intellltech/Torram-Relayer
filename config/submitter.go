package config

import (
	"errors"
	"time"
)

const (
	// Default values for Submitter configuration
	DefaultBufferSize             = 100  // Number of entries in the buffer
	DefaultResubmitFeeMultiplier  = 1.0  // Multiplier for bumped fees
	DefaultPollingIntervalSeconds = 60   // Polling interval (in seconds)
	DefaultResendIntervalSeconds  = 1800 // Resend interval (30 minutes)
)

// SubmitterConfig defines the configuration for the Submitter component.
type SubmitterConfig struct {
	// NetworkParams defines the Bitcoin network parameters (e.g., mainnet, testnet, simnet, signet).
	NetworkParams string `mapstructure:"network-params"`
	// BufferSize specifies the maximum number of raw checkpoints stored in memory.
	BufferSize uint `mapstructure:"buffer-size"`
	// ResubmitFeeMultiplier is a factor to calculate fee increases for resubmissions.
	ResubmitFeeMultiplier float64 `mapstructure:"resubmit-fee-multiplier"`
	// PollingInterval defines how often (in seconds) Torram messages are checked.
	PollingInterval time.Duration `mapstructure:"polling-interval"`
	// ResendInterval specifies the delay (in seconds) before resubmitting to Bitcoin.
	ResendInterval time.Duration `mapstructure:"resend-interval"`
	// DatabaseConfig holds the configuration for the database backend used by the Submitter.
	DatabaseConfig *DBConfig `mapstructure:"database-config"`
}

// Validate checks the integrity of SubmitterConfig values.
func (cfg *SubmitterConfig) Validate() error {
	// Validate NetworkParams
	if !isValidNetworkParam(cfg.NetworkParams) {
		return errors.New("invalid network-params, must be one of: mainnet, testnet, simnet, signet")
	}

	// Validate BufferSize
	if cfg.BufferSize == 0 {
		return errors.New("buffer-size must be greater than 0")
	}

	// Validate ResubmitFeeMultiplier
	if cfg.ResubmitFeeMultiplier < 1.0 {
		return errors.New("resubmit-fee-multiplier must be at least 1.0")
	}

	// Validate PollingInterval
	if cfg.PollingInterval <= 0 {
		return errors.New("polling-interval must be greater than 0 seconds")
	}

	// Validate ResendInterval
	if cfg.ResendInterval <= 0 {
		return errors.New("resend-interval must be greater than 0 seconds")
	}

	// Validate DatabaseConfig
	if cfg.DatabaseConfig == nil {
		return errors.New("database-config cannot be nil")
	}

	if err := cfg.DatabaseConfig.Validate(); err != nil {
		return err
	}

	return nil
}

// DefaultSubmitterConfig provides a default configuration for the Submitter.
func DefaultSubmitterConfig() SubmitterConfig {
	return SubmitterConfig{
		NetworkParams:         "simnet", // Default to Bitcoin Simnet
		BufferSize:            DefaultBufferSize,
		ResubmitFeeMultiplier: DefaultResubmitFeeMultiplier,
		PollingInterval:       time.Second * DefaultPollingIntervalSeconds,
		ResendInterval:        time.Second * DefaultResendIntervalSeconds,
		DatabaseConfig:        DefaultDBConfig(),
	}
}

// isValidNetworkParam checks if a given Bitcoin network parameter is valid.
func isValidNetworkParam(param string) bool {
	validParams := []string{"mainnet", "testnet", "simnet", "signet"}
	for _, p := range validParams {
		if p == param {
			return true
		}
	}
	return false
}
