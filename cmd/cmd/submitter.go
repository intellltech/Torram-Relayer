package cmd

import (
	"fmt"
	"os"

	"github.com/TopDev113/torram-relayer/config"
	"github.com/TopDev113/torram-relayer/internal/btcclient"
	"github.com/TopDev113/torram-relayer/internal/submitter"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

// GetSubmitterCmd initializes and returns the CLI command for the Submitter.
func GetSubmitterCmd() *cobra.Command {
	var cfgFile string

	// Define the submitter command.
	cmd := &cobra.Command{
		Use:   "submitter",
		Short: "Run the Torram to Bitcoin Submitter",
		Run: func(cmd *cobra.Command, args []string) {
			// Load configuration from file.
			cfg, err := config.New(cfgFile)
			if err != nil {
				fmt.Printf("Failed to load configuration: %v\n", err)
				os.Exit(1)
			}

			// Initialize logger.
			logger, err := cfg.CreateLogger()
			if err != nil {
				fmt.Printf("Failed to create logger: %v\n", err)
				os.Exit(1)
			}
			logger.Info("Submitter configuration loaded successfully")

			// Initialize Bitcoin client.
			btcClient, err := btcclient.NewWallet(&cfg, logger)
			if err != nil {
				logger.Error("Failed to initialize Bitcoin client", err)
				os.Exit(1)
			}
			logger.Info("Bitcoin client initialized")

			// Create submitter instance.
			submitterInstance, err := submitter.New(
				&cfg.Submitter,
				logger,
				btcClient,
				nil,              // Replace with Torram query client when integrated
				sdk.AccAddress{}, // Replace with Torram submitter address when available
				cfg.Common.RetrySleepTime,
				cfg.Common.MaxRetrySleepTime,
				cfg.Common.MaxRetryTimes,
				nil, // Replace with metrics instance when integrated
				cfg.Common.DBBackend,
				cfg.BTC.WalletName,
			)
			if err != nil {
				logger.Error("Failed to create submitter", "error", err)
				os.Exit(1)
			}
			logger.Info("Submitter instance created")

			// Start the submitter.
			logger.Info("Starting the submitter...")
			if err := submitterInstance.Start(); err != nil {
				logger.Error("Failed to start the submitter", "error", err)
				os.Exit(1)
			}

			logger.Info("Submitter started successfully. Running...")

			// Block the process to keep the submitter running.
			select {}
		},
	}

	// Add a flag to specify the config file.
	cmd.Flags().StringVar(&cfgFile, "config", config.DefaultConfigFile(), "Path to the configuration file")

	return cmd
}
