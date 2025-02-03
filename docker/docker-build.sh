#!/bin/bash
set -e  # Exit on error
set -x  # Debugging output

# Define variables
GOWEBPATH="/opt/goweb"
APPDIR="$GOWEBPATH/www"
LOGPATH="$GOWEBPATH/logs"
LOGSHPATH="$APPDIR/logging/logs.sh"
SVCPATH="$APPDIR/logging/weblogs.service"
SVCNAME="weblogs.service"
DOCKNAME="goweb_1"
IMGNAME="goweb"

# Function to create a directory if it doesn't exist
create_dir() {
    local dir="$1"
    if [ ! -d "$dir" ]; then
        mkdir -p "$dir"
        chown goweb:goweb "$dir"
    fi
}

# Ensure the 'goweb' user exists
if ! id "goweb" &>/dev/null; then
    useradd -m -d "$GOWEBPATH" -s /bin/bash goweb
fi

# Create required directories
create_dir "$GOWEBPATH"
create_dir "$LOGPATH"

# Set up application directory
if [ -d "$APPDIR" ]; then
    rm -rf "$APPDIR"
fi
mkdir -p "$APPDIR"
chown goweb:goweb "$APPDIR"

# Copy application files
cp -r ../* "$APPDIR/"
cp "$APPDIR/docker/Dockerfile" "$APPDIR"
chown -R goweb:goweb $APPDIR/*
chmod +x "$LOGSHPATH"

# Build Docker image
docker build --tag "$IMGNAME" "$APPDIR/docker"

# Remove existing container if it exists
if docker ps -a -f name="$DOCKNAME" --format '{{.Names}}' | grep -q "$DOCKNAME"; then
    docker rm -f "$DOCKNAME"
fi

# Run Docker container
docker run -d --name "$DOCKNAME" -p 80:8080 "$IMGNAME:latest"

# Manage systemd service
if ! systemctl is-active --quiet "$SVCNAME"; then
    if [ ! -f "/etc/systemd/system/$SVCNAME" ]; then
        cp "$SVCPATH" /etc/systemd/system/
        systemctl enable "$SVCNAME"
    fi
    systemctl start "$SVCNAME"
else
    systemctl restart "$SVCNAME"
fi
