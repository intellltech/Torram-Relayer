# Define directories
PROTO_DIR := proto
GO_OUT := $(PROTO_DIR)
CONFIG_FILE := config/config.toml

# Define executables
PROTOC := protoc
GO := go
GOTEST := $(GO) test
GOBUILD := $(GO) build
GORUN := $(GO) run
GOTIDY := $(GO) mod tidy
GOVET := $(GO) vet
BIN := relayer

# Default target
all: build

# Build the project
build: 
	$(GOBUILD) -o $(BIN) ./cmd/relayer

# Run the project
run: 
	$(GORUN) ./cmd/relayer --config $(CONFIG_FILE)

# Test the project
test: 
	$(GOTEST) ./... -v

# Clean build artifacts
clean:
	rm -f $(BIN)

# Generate .proto files
proto:
	$(PROTOC) --go_out=$(GO_OUT) --go-grpc_out=$(GO_OUT) $(PROTO_DIR)/*.proto

# Lint and format
lint:
	$(GOVET) ./...
	$(GO) fmt ./...

# Install dependencies
deps:
	$(GOTIDY)

# Help menu
help:
	@echo "Usage:"
	@echo "  make build        Build the relayer binary"
	@echo "  make run          Run the relayer application"
	@echo "  make test         Run all tests"
	@echo "  make clean        Remove build artifacts"
	@echo "  make proto        Compile .proto files"
	@echo "  make lint         Run linters and format code"
	@echo "  make deps         Install dependencies"
	@echo "  make help         Show this help message"
