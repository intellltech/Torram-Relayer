#!/bin/bash
# build_protos.sh

echo "Compiling .proto files..."

PROTO_DIR="./proto"
OUT_DIR="./proto"

if [ ! -d "$OUT_DIR" ]; then
    mkdir -p "$OUT_DIR"
fi

protoc --go_out=$OUT_DIR --go-grpc_out=$OUT_DIR $PROTO_DIR/*.proto

if [ $? -eq 0 ]; then
    echo "Protobuf compilation complete!"
else
    echo "Error: Failed to compile .proto files."
    exit 1
fi
