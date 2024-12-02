package submitter

import (
	"log"
	"time"

	"github.com/TopDev113/torram-relayer/config"
	"github.com/TopDev113/torram-relayer/internal/btcclient"
	"github.com/TorramLabs-Team/TorramChain/x/torram/types"
)

type Submitter struct {
	cfg       *config.Config
	logger    *log.Logger
	btcWallet *btcclient.Wallet
}

// NewSubmitter initializes a new submitter instance
func NewSubmitter(cfg *config.Config, logger *log.Logger, btcWallet *btcclient.Wallet) *Submitter {
	return &Submitter{
		cfg:       cfg,
		logger:    logger,
		btcWallet: btcWallet,
	}
}

// Start begins the submitter service for relaying messages
func (s *Submitter) Start() error {
	s.logger.Println("Starting the submitter service...")

	// Example: Simulating a periodic check for Torram-to-Bitcoin messages
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for range ticker.C {
		s.logger.Println("Checking for Torram messages...")

		// Simulate Torram message handling
		msg := &types.TransactionData{ID: "example_tx"}
		err := s.submitToBitcoin(msg)
		if err != nil {
			s.logger.Printf("Failed to submit transaction to Bitcoin: %v", err)
		} else {
			s.logger.Printf("Successfully submitted transaction: %s", msg.ID)
		}
	}

	return nil
}

// submitToBitcoin handles the process of submitting Torram transactions to Bitcoin
func (s *Submitter) submitToBitcoin(txData *types.TransactionData) error {
	// Simulate converting Torram transaction to Bitcoin format
	btcTx, err := s.btcWallet.CreateTransaction(txData)
	if err != nil {
		return err
	}

	// Send transaction to Bitcoin network
	txHash, err := s.btcWallet.SendTransaction(btcTx)
	if err != nil {
		return err
	}

	// Log success
	s.logger.Printf("Bitcoin transaction submitted: %s", txHash)
	return nil
}
