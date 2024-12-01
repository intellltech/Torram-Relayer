# Step 1: Build the Go binary
FROM golang:1.20 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the entire project to the container
COPY . .

# Build the Go binary
RUN GOOS=linux GOARCH=amd64 go build -o relayer .

# Step 2: Create a smaller runtime image
FROM alpine:latest

# Install necessary dependencies (if any)
RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Go binary from the builder stage
COPY --from=builder /app/relayer .

# Copy the config file
COPY config/config.toml ./config/config.toml

# Expose the port your application will run on (example)
EXPOSE 8080

# Command to run the binary
CMD ["./relayer"]
