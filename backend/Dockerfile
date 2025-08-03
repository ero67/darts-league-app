# Build stage
FROM golang:1.23-alpine AS builder

# Install git and ca-certificates (needed for go modules)
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Create appuser
RUN adduser -D -g '' appuser

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download
RUN go mod verify

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o main ./cmd/server

# Final stage
FROM scratch

# Import from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy our static executable
COPY --from=builder /app/main /app/main

# Use an unprivileged user
USER appuser

# Expose port
EXPOSE 8080

# Run the binary
ENTRYPOINT ["/app/main"]