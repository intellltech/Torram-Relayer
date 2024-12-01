package torram

import (
	"context"
	"log"

	sdktx "github.com/cosmos/cosmos-sdk/types/tx"
	"google.golang.org/grpc"
)

func BroadcastTransaction(grpcEndpoint string, txBytes []byte) error {
	conn, err := grpc.Dial(grpcEndpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	txService := sdktx.NewServiceClient(conn)

	resp, err := txService.BroadcastTx(context.Background(), &sdktx.BroadcastTxRequest{
		Mode:    sdktx.BroadcastMode_BROADCAST_MODE_SYNC,
		TxBytes: txBytes,
	})
	if err != nil {
		return err
	}

	log.Printf("Transaction Response: %v", resp)
	return nil
}
