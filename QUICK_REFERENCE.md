# Network Configuration Manager - Quick Reference

## Installation

### Windows
```powershell
.\install.ps1
```

### Linux/macOS
```bash
bash install.sh
```

## Quick Commands

### VLAN Commands
```bash
# Create VLAN
network-config vlan create --id 10 --name "Production"

# List all VLANs
network-config vlan list

# Delete VLAN
network-config vlan delete --id 10
```

### VPN Commands
```bash
# Create L2TP VPN
network-config vpn l2tp --action create --name "office"

# Connect to VPN
network-config vpn l2tp --action connect --name "office"

# Create PPTP VPN
network-config vpn pptp --action create --name "backup"

# Create SSTP VPN
network-config vpn sstp --action create --name "secure"

# Start OpenVPN
network-config vpn openvpn --action start --config "config.ovpn"
```

### Firewall Commands
```bash
# Add rule
network-config firewall rule --action add --rule "allow tcp port 80"

# List rules
network-config firewall list

# Check status
network-config firewall status
```

### QoS Commands
```bash
# Create policy
network-config qos policy --name "high" --bandwidth "100Mbps"

# Set limit
network-config qos limit --interface eth0 --limit "50Mbps"

# Set priority
network-config qos priority --class voip --priority 7
```

### Monitoring Commands
```bash
# Monitor traffic
network-config monitor traffic --interface eth0

# Get statistics
network-config monitor stats

# Configure alerts
network-config monitor alerts --threshold "100Mbps"
```

## Common Tasks

### Setup Production Network
```bash
# Create VLANs
network-config vlan create --id 1 --name "Management"
network-config vlan create --id 10 --name "Production"
network-config vlan create --id 100 --name "DMZ"

# Configure Firewall
network-config firewall rule --action add --rule "allow tcp port 80"
network-config firewall rule --action add --rule "allow tcp port 443"
network-config firewall rule --action add --rule "allow tcp port 22"

# Setup QoS
network-config qos policy --name "critical" --bandwidth "200Mbps"
network-config qos priority --class critical --priority 7

# Start Monitoring
network-config monitor traffic --interface eth0
```

### Secure Remote Access
```bash
# Create OpenVPN
network-config vpn openvpn --action start --config "vpn.conf"

# Fallback L2TP
network-config vpn l2tp --action create --name "backup-vpn"

# Enable Security Rules
network-config firewall rule --action add --rule "allow tcp port 443"
network-config firewall rule --action add --rule "allow udp port 1194"
```

### Monitor Network Health
```bash
# Setup alerts
network-config monitor alerts --threshold "150Mbps"

# Check real-time stats
network-config monitor traffic --interface eth0

# View network statistics
network-config monitor stats
```

## Configuration Examples

### network.yml Format
```yaml
vlans:
  - id: "10"
    name: "Production"
    status: "active"

vpn:
  - name: "office-vpn"
    type: "L2TP"
    server: "vpn.example.com"

firewall:
  - direction: "Inbound"
    action: "Allow"
    protocol: "TCP"
    port: "80"

qos:
  - name: "high-priority"
    bandwidth: "100Mbps"
    priority: 7

monitoring:
  enabled: true
  interval: 5
```

## Flags Reference

### Global Flags
```
--config FILE      Configuration file
--loglevel LEVEL   Log level (debug, info, warn, error)
--verbose          Verbose output
--help            Show help
```

### VLAN Flags
```
--id ID             VLAN ID
--name NAME         VLAN name
--description DESC  Description
```

### VPN Flags
```
--action ACTION     Action (create, connect, disconnect)
--name NAME         Connection name
--config FILE       Config file path
```

### Firewall Flags
```
--action ACTION     Action (add, remove, modify)
--rule RULE         Rule definition
```

### QoS Flags
```
--name NAME         Policy name
--bandwidth BW      Bandwidth (e.g., 100Mbps)
--interface IF      Interface name
--priority LEVEL    Priority (0-7)
```

### Monitor Flags
```
--interface IF      Interface to monitor
--threshold THR     Alert threshold
```

## Troubleshooting

### VPN Connection Failed
```bash
# Check VPN status
network-config vpn l2tp --action connect --name "office"

# Try fallback
network-config vpn pptp --action connect --name "office"
```

### Firewall Issues
```bash
# List all rules
network-config firewall list

# Check status
network-config firewall status
```

### Traffic Monitoring
```bash
# Monitor specific interface
network-config monitor traffic --interface eth0

# Check overall stats
network-config monitor stats
```

## Performance Tips

1. **VLAN Design**
   - Minimize inter-VLAN traffic
   - Group related services

2. **VPN Optimization**
   - Use OpenVPN for best performance
   - Enable compression if needed

3. **Firewall Rules**
   - Order rules by frequency
   - Use groups for multiple ports

4. **QoS Settings**
   - Match real bandwidth
   - Adjust priorities regularly

5. **Monitoring**
   - Set reasonable alert thresholds
   - Review logs periodically

## Getting Help

```bash
# General help
network-config --help

# Command-specific help
network-config vlan --help
network-config vpn --help
network-config firewall --help
network-config qos --help
network-config monitor --help
```

## File Locations

- **Config**: `./config/` or `/etc/network-config/`
- **Logs**: `./logs/` or `/var/log/`
- **Examples**: `./examples/`
- **Binary**: `./network-config` or `./network-config.exe`

## Keyboard Shortcuts

- `Ctrl+C` - Stop monitoring
- `Ctrl+D` - Exit application
- `Tab` - Auto-complete in interactive mode

## Common Aliases

Create shortcuts for frequent commands:

### Linux/macOS
```bash
alias nc-vlan='network-config vlan'
alias nc-vpn='network-config vpn'
alias nc-fw='network-config firewall'
alias nc-mon='network-config monitor'
```

### Windows (PowerShell)
```powershell
Set-Alias -Name nc-vlan -Value ".\network-config.exe vlan"
Set-Alias -Name nc-vpn -Value ".\network-config.exe vpn"
Set-Alias -Name nc-fw -Value ".\network-config.exe firewall"
Set-Alias -Name nc-mon -Value ".\network-config.exe monitor"
```

## Links

- **Documentation**: README.md
- **Usage Guide**: USAGE_GUIDE.md
- **API Docs**: API_DOCUMENTATION.md
- **Examples**: examples/network.yml
- **Advanced Setup**: ADVANCED_CONFIG.md
