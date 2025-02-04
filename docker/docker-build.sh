#!/bin/bash
set -e  # Exit on error

# Get root directory of project
SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

# Docker variables
IMAGE_NAME="andrew-pineiro/goweb:latest"
DOCKER_NAME="goweb_1"

# Logging Services
SVC_NAME="weblogs.service"
SVC_FILEPATH="$PROJECT_ROOT/logging"

# Move Dockerfile to src/ for compiling
cp $SCRIPT_DIR/Dockerfile $PROJECT_ROOT/src
pushd $PROJECT_ROOT/src

echo "Building Docker image..."
docker build -t $IMAGE_NAME .
echo "Docker build complete!"

rm Dockerfile
popd

# Remove existing container if it exists
if docker ps -a -f name="$DOCKER_NAME" --format '{{.Names}}' | grep -q "$DOCKER_NAME"; then
    docker rm -f "$DOCKER_NAME"
fi

# Run Docker container
docker run -d --name "$DOCKER_NAME" -p 80:8080 "$IMAGE_NAME"

# Manage systemd service
if ! systemctl is-active --quiet "$SVC_NAME"; then
    if [ ! -f "/etc/systemd/system/$SVC_NAME" ]; then
        cp "$SVC_FILEPATH" /etc/systemd/system/
        systemctl enable "$SVC_NAME"
    fi
    systemctl start "$SVC_NAME"
else
    systemctl restart "$SVC_NAME"
fi