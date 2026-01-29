# Dockerfile for Network Configuration Manager

FROM golang:1.21-alpine AS builder

WORKDIR /app

# Copy source files
COPY . .

# Build the application
RUN go mod download && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o network-config main.go

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/network-config .

# Copy example configs
COPY examples/ ./examples/

# Create log directory
RUN mkdir -p /var/log

VOLUME ["/etc/network-config", "/var/log"]

EXPOSE 8080 8443

ENTRYPOINT ["./network-config"]
CMD ["--help"]
