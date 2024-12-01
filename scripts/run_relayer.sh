#!/bin/bash
# run_relayer.sh

echo "Starting the relayer..."

if [ -f "./relayer" ]; then
    ./relayer
else
    echo "Error: Relayer binary not found. Please build it first."
    exit 1
fi
