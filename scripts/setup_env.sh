#!/bin/bash
# setup_env.sh

echo "Setting up environment variables for the relayer..."

export TORRAM_RPC_ENDPOINT="http://localhost:26657"
export TORRAM_GRPC_ENDPOINT="localhost:9090"
export BITCOIN_RPC_ENDPOINT="http://localhost:8332"
export BITCOIN_USER="bitcoinrpc"
export BITCOIN_PASSWORD="your_password"

echo "Environment variables set:"
echo "TORRAM_RPC_ENDPOINT=$TORRAM_RPC_ENDPOINT"
echo "TORRAM_GRPC_ENDPOINT=$TORRAM_GRPC_ENDPOINT"
echo "BITCOIN_RPC_ENDPOINT=$BITCOIN_RPC_ENDPOINT"
echo "BITCOIN_USER=$BITCOIN_USER"
echo "BITCOIN_PASSWORD=[hidden]"

echo "Installing Go dependencies..."
go mod tidy

echo "Setup complete!"
