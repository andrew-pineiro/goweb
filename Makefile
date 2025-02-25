BINARY_NAME=goweb
BUILD_DIR=bin
OLD_DIR=$(BUILD_DIR)/.old
SRC_DIR=src

# Default to your system's OS and architecture
GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

.PHONY: build clean build-all

# Single build with specified GOOS/GOARCH
build:
	@mkdir -p $(BUILD_DIR)
	@echo "building goweb application..."
	@CGO_ENABLED=$$( [ "$(GOOS)" = "linux" ] && [ "$(GOARCH)" = "amd64" ] && echo 1 || echo 0 ) \
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -C $(SRC_DIR) -o ../$(BUILD_DIR)/$(BINARY_NAME)-$(GOOS)-$(GOARCH)
	@echo "success"

# Build for multiple platforms
build-all: clean
	mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -C $(SRC_DIR) -o ../$(BUILD_DIR)/$(BINARY_NAME)-linux-amd64
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -C $(SRC_DIR) -o ../$(BUILD_DIR)/$(BINARY_NAME)-windows-amd64.exe
	GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -C $(SRC_DIR) -o ../$(BUILD_DIR)/$(BINARY_NAME)-darwin-amd64
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -C $(SRC_DIR) -o ../$(BUILD_DIR)/$(BINARY_NAME)-linux-arm64
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -C $(SRC_DIR) -o ../$(BUILD_DIR)/$(BINARY_NAME)-darwin-arm64 

# Clean up binaries
clean:
	rm -rf $(OLD_DIR)
	mkdir -p $(OLD_DIR)
	mv $(BUILD_DIR)/* $(OLD_DIR)/ 2>/dev/null || true
