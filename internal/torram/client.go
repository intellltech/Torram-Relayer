package torram

import (
	"context"

	torramproto "github.com/TorramLabs-Team/TorramChain/x/torram/types"
	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client torramproto.QueryClient
}

func NewClient(grpcEndpoint string) (*Client, error) {
	conn, err := grpc.Dial(grpcEndpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		client: torramproto.NewQueryClient(conn),
	}, nil
}

func (c *Client) QueryGame(ctx context.Context, gameId string) (*torramproto.Game, error) {
	response, err := c.client.Query(ctx, &torramproto.QueryGetGameRequest{Id: gameId}) // Adjust to correct method name
	if err != nil {
		return nil, err
	}
	return response.GetGame(), nil // Adjust to return the correct type
}

func (c *Client) Close() error {
	return c.conn.Close()
}

// Method to get the GRPC endpoint from the Client.
func (c *Client) GetRPCEndpoint() string {
	return c.conn.Target() // Returns the GRPC connection's target (endpoint)
}
