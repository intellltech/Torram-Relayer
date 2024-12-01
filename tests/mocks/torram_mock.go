package mocks

// import (
// 	"context"

// 	"github.com/TopDev113/torram-relayer/proto/torram_relayer"
// )

// type TorramMock struct{}

// func (t *TorramMock) StreamTorramEvents(req *torram_relayer.StreamEventsRequest, stream torram_relayer.TorramRelayerService_StreamTorramEventsServer) error {
// 	mockEvent := &torram_relayer.TorramEventResponse{
// 		EventId:   "1234",
// 		EventType: "transfer",
// 		Payload:   []byte("mock-data"),
// 	}
// 	return stream.Send(mockEvent)
// }

// func (t *TorramMock) SubmitBitcoinTransaction(ctx context.Context, req *torram_relayer.BitcoinTransactionRequest) (*torram_relayer.TransactionResponse, error) {
// 	return &torram_relayer.TransactionResponse{
// 		Success: true,
// 		Message: "Transaction submitted successfully",
// 	}, nil
// }
