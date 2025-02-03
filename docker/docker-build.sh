#!/bin/sh
set -e  # Exit on error

IMAGE_NAME="andrew-pineiro/goweb:latest"

cp Dockerfile ../src

echo "Building Docker image..."
pushd ../src
docker build -t $IMAGE_NAME .

echo "Docker build complete!"
rm Dockerfile
popd