// In internal/torram/events.go

package torram

import (
	"context"
	"log"
	"time"

	tmclient "github.com/tendermint/tendermint/rpc/client"
)

func SubscribeToEvents(ctx context.Context, rpcEndpoint string, eventType string, handler func(event interface{})) error {
	client, err := tmclient.NewHTTP(rpcEndpoint, "/websocket")
	if err != nil {
		return err
	}

	// Start the client
	if err := client.Start(); err != nil {
		return err
	}

	query := "tm.event = '" + eventType + "'" // Example event filter
	subs, err := client.Subscribe(ctx, "subscriber", query)
	if err != nil {
		return err
	}

	// Listen for events
	for {
		select {
		case event := <-subs:
			handler(event)
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(30 * time.Second): // Heartbeat to ensure connection is alive
			log.Println("Waiting for events...")
		}
	}
}
