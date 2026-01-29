# Network Configuration Manager - Windows Setup

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Network Config Manager - Windows Setup"
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# Check if Go is installed
Write-Host "Checking for Go installation..." -ForegroundColor Yellow

if (!(Get-Command go -ErrorAction SilentlyContinue)) {
    Write-Host "Error: Go is not installed. Please install Go 1.21 or later." -ForegroundColor Red
    Write-Host "Download from: https://golang.org/dl/" -ForegroundColor Green
    exit 1
}

Write-Host "✓ Go is installed" -ForegroundColor Green
Write-Host ""

# Get Go version
$goVersion = go version
Write-Host "Go version: $goVersion" -ForegroundColor Green
Write-Host ""

# Download dependencies
Write-Host "Installing dependencies..." -ForegroundColor Yellow
go mod download

if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ Dependencies installed" -ForegroundColor Green
} else {
    Write-Host "✗ Failed to install dependencies" -ForegroundColor Red
    exit 1
}
Write-Host ""

# Build the project
Write-Host "Building project..." -ForegroundColor Yellow
go build -o network-config.exe main.go

if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ Project built successfully" -ForegroundColor Green
    Write-Host "Binary: network-config.exe" -ForegroundColor Green
} else {
    Write-Host "✗ Build failed" -ForegroundColor Red
    exit 1
}
Write-Host ""

# Run tests
Write-Host "Running tests..." -ForegroundColor Yellow
go test ./tests -v

if ($LASTEXITCODE -eq 0) {
    Write-Host "✓ All tests passed" -ForegroundColor Green
} else {
    Write-Host "⚠ Some tests may have issues" -ForegroundColor Yellow
}
Write-Host ""

# Display help
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Installation Complete!" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "Commands:" -ForegroundColor Yellow
Write-Host "  .\network-config.exe --help           Show help" -ForegroundColor White
Write-Host "  .\network-config.exe vlan create      Create VLAN" -ForegroundColor White
Write-Host "  .\network-config.exe vpn l2tp         Configure L2TP VPN" -ForegroundColor White
Write-Host "  .\network-config.exe firewall list    List firewall rules" -ForegroundColor White
Write-Host "  .\network-config.exe qos policy       Create QoS policy" -ForegroundColor White
Write-Host "  .\network-config.exe monitor traffic  Monitor traffic" -ForegroundColor White
Write-Host ""
