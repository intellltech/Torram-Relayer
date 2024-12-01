#!/bin/bash
# test_relayer.sh

echo "Running relayer tests..."

# Unit tests
echo "Running unit tests..."
go test ./... -v

# Integration tests
echo "Running integration tests..."
go test ./internal/relay/... -v

if [ $? -eq 0 ]; then
    echo "All tests passed!"
else
    echo "Tests failed. Please check the logs."
    exit 1
fi
