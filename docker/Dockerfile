# Build Stage
FROM golang:1.19-alpine AS builder

# Set up working directory
WORKDIR /app

# Install necessary dependencies
RUN apk add --no-cache git

# Copy go mod and sum files separately to leverage caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o goweb

# Final Stage (Minimal Image)
FROM scratch

# Set working directory
WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=builder /app/goweb .

# Add a non-root user for security
USER 1001

# Expose port 8080
EXPOSE 8080

# Define the entrypoint
ENTRYPOINT ["./goweb"]
