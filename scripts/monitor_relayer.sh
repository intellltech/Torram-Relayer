#!/bin/bash
# monitor_relayer.sh

LOG_FILE="./logs/relayer.log"

if [ ! -f "$LOG_FILE" ]; then
    echo "Error: Log file not found at $LOG_FILE"
    exit 1
fi

echo "Tailing relayer logs..."
tail -f "$LOG_FILE"
