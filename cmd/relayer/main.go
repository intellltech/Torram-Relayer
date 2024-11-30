package main

import (
	"time"

	"github.com/TopDev113/torram-relayer/config"
	"github.com/TopDev113/torram-relayer/internal/core"
	"github.com/TopDev113/torram-relayer/internal/utils"
	"go.uber.org/zap"
)

func main() {
	logger := utils.NewLogger()
	defer logger.Sync()

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Failed to load configuration", zap.Error(err))
	}

	// Initialize Bitcoin and Torram clients
	btcClient := core.NewBitcoinHandler(cfg.Bitcoin)
	torramClient := core.NewTorramHandler(cfg.Torram)

	// Create the relayer
	relayer := core.NewRelayer(btcClient, torramClient, logger, time.Duration(cfg.PollIntervalSeconds)*time.Second)

	// Start the relayer
	relayer.Start()
}
