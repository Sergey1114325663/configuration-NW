#!/bin/bash

# Network Configuration Manager - Installation and Setup Script

echo "================================"
echo "Network Config Manager Setup"
echo "================================"
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go 1.21 or later."
    exit 1
fi

echo "✓ Go is installed"
echo ""

# Download dependencies
echo "Installing dependencies..."
go mod download

if [ $? -eq 0 ]; then
    echo "✓ Dependencies installed"
else
    echo "✗ Failed to install dependencies"
    exit 1
fi
echo ""

# Build the project
echo "Building project..."
go build -o network-config main.go

if [ $? -eq 0 ]; then
    echo "✓ Project built successfully"
else
    echo "✗ Build failed"
    exit 1
fi
echo ""

# Run tests
echo "Running tests..."
go test ./tests -v

if [ $? -eq 0 ]; then
    echo "✓ All tests passed"
else
    echo "✗ Some tests failed"
fi
echo ""

# Display help
echo "================================"
echo "Installation Complete!"
echo "================================"
echo ""
echo "Run: ./network-config --help"
echo "for available commands"
echo ""
