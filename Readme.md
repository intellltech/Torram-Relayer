# Torram Relayer

A relayer for connecting the Torram Chain and Bitcoin, designed to forward events between both networks. This project is intended for use in decentralized finance (DeFi) applications, blockchain interoperability, and cross-chain operations.

## Features
- **Torram Chain** integration for processing events.
- **Bitcoin** integration via JSON-RPC for event handling.
- Configuration via TOML files.
- Built using Go, with Protocol Buffers (gRPC) support for inter-process communication.

## Project Structure
- `internal/`: Contains core application code.
  - `btcclient/`: Bitcoin client code.
  - `torram/`: Torram client code.
  - `submitter/`: submitter relayer functionality.
- `proto/`: Protocol Buffers definitions.
- `cmd/`: All commands for Relayer modules.
- `tests/`: Test relayer functionality between Torram-Chain and Bitcoin node.
- `config/`: Configuration files (e.g., `config.toml`).
- `scripts/`: Utility scripts.
- `Makefile`: Makefile for building and running the project.
- `Dockerfile`: Docker configuration for containerizing the application.

## Installation

### Prerequisites
- Go 1.20 or higher
- Docker (optional for containerized environment)
- A working Torram Chain node and Bitcoin RPC node.

### Clone the Repository
```bash
git clone https://github.com/TopDev113/torram-relayer.git
cd torram-relayer
```

### instsall Dependencies
Run the following command to install necessary Go dependencies:
```bash
go mod tidy
```

### instsall Dependencies
You can build the relayer with the following command:
```bash
go build -o relayer .
```
Alternatively, you can use the provided Makefile:
```bash
make build
```
## Configuration
The relayer expects a config/config.toml file. Here’s an example configuration:
```toml
[Torram]
GRPCEndpoint = "localhost:9090"

[Bitcoin]
RPCEndpoint = "http://localhost:8332"
User = "bitcoin_user"
Password = "bitcoin_password"
```

## Running the Relayer

After building the application, you can run the relayer with:
```
./relayer
```
Alternatively, you can use Docker:
1. **Build the Docker image** 
    ```bash
    docker build -t torram-relayer .
    ```
2. **Run the Docker container** 
    ```bash
    docker run -d -p 8080:8080 torram-relayer
    ```

## Usage
Once the relayer is running, it will continuously monitor the Torram Chain for events and relay relevant data to the Bitcoin blockchain (or vice versa). For more specific functionality, you can adjust the event subscriptions and processing logic in the code.

## Developing
To modify or extend the relayer functionality:

1.  **Add new event types**: Modify *torram/* and *btcclient/* directories to add more blockchain-specific functionality.
2.  **Add new services or APIs**: Add or Modify the **relayer module** directory in *internal/* like *submitter/* to expose more services.
3.  **Protocol Buffers**: If you modify or add .proto files, don’t forget to regenerate the Go code by running:
    ```bash
    make proto
    ```

## Testing
To run tests ofr the relayer:
```bash
make test
```
To test the individual components, run the following:
```bash
go test ./internal/...
```

## Docker Compose (Optional)
You can use Docker Compose to run the relayer along with other services (e.g., Torram and Bitcoin nodes). Here is an example docker-compose.yml:

```yaml
version: '3.8'

services:
  relayer:
    build: .
    ports:
      - "8080:8080"
    volumes:
      - ./config:/app/config
    environment:
      - CONFIG_FILE=/app/config/config.toml
```
To start the containers:
```bash
docker-compose up --build
```