package core

import (
	"time"

	"go.uber.org/zap"
)

type Relayer struct {
	BTCClient    *BitcoinHandler
	TorramClient *TorramHandler
	Logger       *zap.SugaredLogger
	PollInterval time.Duration
	Quit         chan struct{}
}

func NewRelayer(btcClient *BitcoinHandler, torramClient *TorramHandler, logger *zap.Logger, pollInterval time.Duration) *Relayer {
	return &Relayer{
		BTCClient:    btcClient,
		TorramClient: torramClient,
		Logger:       logger.Sugar(),
		PollInterval: pollInterval,
		Quit:         make(chan struct{}),
	}
}

func (r *Relayer) Start() {
	r.Logger.Info("Starting relayer...")
	ticker := time.NewTicker(r.PollInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := r.syncData(); err != nil {
				r.Logger.Errorf("Sync failed: %v", err)
			}
		case <-r.Quit:
			r.Logger.Info("Stopping relayer...")
			return
		}
	}
}

func (r *Relayer) syncData() error {
	r.Logger.Info("Relaying data between Bitcoin and Torram...")

	// Fetch Bitcoin checkpoints
	btcData, err := r.BTCClient.FetchCheckpoints()
	if err != nil {
		return err
	}

	// Submit to Torram
	if err := r.TorramClient.SubmitCheckpoints(btcData); err != nil {
		return err
	}

	// Fetch Torram checkpoints
	torramData, err := r.TorramClient.FetchCheckpoints()
	if err != nil {
		return err
	}

	// Submit to Bitcoin
	if err := r.BTCClient.SubmitCheckpoints(torramData); err != nil {
		return err
	}

	return nil
}
