package integration

// import (
// 	"context"
// 	"testing"

// 	"github.com/TopDev113/torram-relayer/proto/torram_relayer"

// 	"github.com/stretchr/testify/assert"
// 	"google.golang.org/grpc"
// )

// func TestTorramToBitcoinFlow(t *testing.T) {
// 	// Set up Torram mock server
// 	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure())
// 	assert.NoError(t, err, "GRPC connection to Torram should not error")

// 	client := torram_relayer.NewTorramRelayerServiceClient(conn)

// 	// Mock event streaming
// 	stream, err := client.StreamTorramEvents(context.Background(), &torram_relayer.StreamEventsRequest{
// 		Filter: "tx.height > 0",
// 	})
// 	assert.NoError(t, err, "StreamEventsRequest should not error")

// 	event, err := stream.Recv()
// 	assert.NoError(t, err, "Receiving event should not error")
// 	assert.NotNil(t, event, "Received event should not be nil")

// 	// Simulate Bitcoin submission
// 	response, err := client.SubmitBitcoinTransaction(context.Background(), &torram_relayer.BitcoinTransactionRequest{
// 		TxId:      "abcd1234",
// 		Recipient: "torram_address",
// 		Amount:    0.1,
// 	})
// 	assert.NoError(t, err, "SubmitBitcoinTransaction should not error")
// 	assert.True(t, response.Success, "Transaction submission should succeed")
// }
