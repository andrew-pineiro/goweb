#!/bin/sh
set -e  # Exit on error

# Set Go environment variables
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

# Build the Go application
echo "Building Go application..."
go build -ldflags="-w -s" -o goweb

# Build the Docker image
echo "Building Docker image..."
docker build -t andrew-pineiro/goweb:latest .

echo "Docker build complete!"
