#!/bin/sh
set -e  # Exit on error

IMAGE_NAME="andrew-pineiro/goweb:latest"
CONTAINER_NAME="goweb-builder"

echo "Building Go application inside Docker..."

# Step 1: Use a temporary container to build the Go application
docker run --rm \
    -v "$(pwd)":/app \
    -w /app \
    golang:1.19-alpine \
    sh -c "apk add --no-cache git && go mod download && go build -ldflags='-w -s' -o goweb"

echo "Go application built successfully!"

# Step 2: Build the final Docker image
echo "Building Docker image..."
docker build -t $IMAGE_NAME .

echo "Docker build complete!"
