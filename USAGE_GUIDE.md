# Network Configuration Manager - Usage Guide

## Quick Start

### 1. Build the project
```bash
go mod download
go build -o network-config main.go
```

### 2. View available commands
```bash
./network-config --help
```

## VLAN Management Examples

### Create VLANs
```bash
# Create production VLAN
./network-config vlan create --id 10 --name "Production" --description "Production Network"

# Create development VLAN
./network-config vlan create --id 20 --name "Development" --description "Dev Network"

# Create guest VLAN
./network-config vlan create --id 30 --name "Guest" --description "Guest Network"
```

### List VLANs
```bash
./network-config vlan list
```

### Delete VLAN
```bash
./network-config vlan delete --id 30
```

## VPN Management Examples

### L2TP VPN
```bash
# Create L2TP connection
./network-config vpn l2tp --action create --name "office-vpn"

# Connect to VPN
./network-config vpn l2tp --action connect --name "office-vpn"

# Disconnect
./network-config vpn l2tp --action disconnect --name "office-vpn"
```

### PPTP VPN
```bash
./network-config vpn pptp --action create --name "remote-vpn"
./network-config vpn pptp --action connect --name "remote-vpn"
```

### SSTP VPN (Secure Socket Tunneling Protocol)
```bash
./network-config vpn sstp --action create --name "secure-vpn"
./network-config vpn sstp --action connect --name "secure-vpn"
```

### OpenVPN
```bash
# Start OpenVPN with config
./network-config vpn openvpn --action start --config "/path/to/config.ovpn"

# Stop OpenVPN
./network-config vpn openvpn --action stop --config "/path/to/config.ovpn"

# Restart OpenVPN
./network-config vpn openvpn --action restart --config "/path/to/config.ovpn"
```

## Firewall Management Examples

### Add Rules
```bash
# Allow HTTP traffic
./network-config firewall rule --action add --rule "allow tcp port 80"

# Allow HTTPS traffic
./network-config firewall rule --action add --rule "allow tcp port 443"

# Block DNS from external
./network-config firewall rule --action add --rule "deny udp port 53"
```

### View Firewall Status
```bash
./network-config firewall status
```

### List Rules
```bash
./network-config firewall list
```

## QoS Configuration Examples

### Create QoS Policies
```bash
# High priority policy
./network-config qos policy --name "high-priority" --bandwidth "100Mbps"

# Normal priority policy
./network-config qos policy --name "normal-priority" --bandwidth "50Mbps"

# Low priority policy
./network-config qos policy --name "low-priority" --bandwidth "10Mbps"
```

### Set Bandwidth Limits
```bash
# Limit eth0 to 100Mbps
./network-config qos limit --interface eth0 --limit "100Mbps"

# Limit eth1 to 50Mbps
./network-config qos limit --interface eth1 --limit "50Mbps"
```

### Set Traffic Priority
```bash
# VoIP traffic - highest priority
./network-config qos priority --class voip --priority 7

# Video streaming - high priority
./network-config qos priority --class video --priority 6

# Web browsing - normal priority
./network-config qos priority --class web --priority 4

# File transfer - low priority
./network-config qos priority --class file-transfer --priority 1
```

## Traffic Monitoring Examples

### Monitor Interface Traffic
```bash
# Monitor eth0
./network-config monitor traffic --interface eth0

# Monitor eth1
./network-config monitor traffic --interface eth1
```

### Get Network Statistics
```bash
./network-config monitor stats
```

### Configure Traffic Alerts
```bash
# Alert when bandwidth exceeds 100Mbps
./network-config monitor alerts --threshold "100Mbps"

# Alert at 80% utilization
./network-config monitor alerts --threshold "80%"
```

## Configuration File Usage

### Load configuration from YAML
```bash
# Create config file
cat > network-config.yml << EOF
vlans:
  - id: "1"
    name: "Management"
    status: "active"
  - id: "10"
    name: "Production"
    status: "active"

vpn:
  - name: "office-vpn"
    type: "L2TP"
    server: "vpn.example.com"
    enabled: true

firewall:
  - direction: "Inbound"
    action: "Allow"
    protocol: "TCP"
    port: "80"

qos:
  - name: "high-priority"
    bandwidth: "100Mbps"
    priority: 7
EOF
```

## Advanced Configuration

### Multi-VLAN Setup
```bash
# Create management VLAN
./network-config vlan create --id 1 --name "Management"

# Create production VLANs
./network-config vlan create --id 10 --name "Production-A"
./network-config vlan create --id 11 --name "Production-B"

# Create segregated networks
./network-config vlan create --id 100 --name "DMZ"
./network-config vlan create --id 200 --name "Guest"
```

### Secured VPN Setup
```bash
# Create secure OpenVPN connection
./network-config vpn openvpn --action start \
  --config "/etc/openvpn/secure.conf"

# Fallback to L2TP if OpenVPN fails
./network-config vpn l2tp --action connect --name "office-vpn"
```

### Comprehensive Firewall Configuration
```bash
# Allow essential services
./network-config firewall rule --action add --rule "allow tcp port 22"   # SSH
./network-config firewall rule --action add --rule "allow tcp port 80"   # HTTP
./network-config firewall rule --action add --rule "allow tcp port 443"  # HTTPS
./network-config firewall rule --action add --rule "allow tcp port 53"   # DNS

# Deny unnecessary services
./network-config firewall rule --action add --rule "deny tcp port 23"    # Telnet
./network-config firewall rule --action add --rule "deny tcp port 3389"  # RDP (if not needed)
```

### QoS with Multiple Priorities
```bash
# VoIP - Highest priority
./network-config qos policy --name "voip" --bandwidth "20Mbps"
./network-config qos priority --class voip --priority 7

# Video conferencing - High priority
./network-config qos policy --name "video-conf" --bandwidth "50Mbps"
./network-config qos priority --class video-conf --priority 6

# Web services - Normal priority
./network-config qos policy --name "web" --bandwidth "100Mbps"
./network-config qos priority --class web --priority 4

# Downloads - Low priority
./network-config qos policy --name "downloads" --bandwidth "30Mbps"
./network-config qos priority --class downloads --priority 1
```

## Performance Tuning

### Monitor and Optimize
```bash
# Check current traffic
./network-config monitor traffic --interface eth0

# Get statistics
./network-config monitor stats

# Set performance alerts
./network-config monitor alerts --threshold "150Mbps"
```

### Optimization Tips
1. Place VLANs strategically to minimize inter-VLAN traffic
2. Use firewall rules in order of frequency (most common first)
3. Monitor QoS policies regularly and adjust priorities
4. Set up alerts for abnormal traffic patterns
5. Keep firewall rules updated and remove obsolete rules

## Troubleshooting

### Check VPN Connection
```bash
./network-config vpn l2tp --action connect --name "office-vpn"
```

### Verify Firewall Rules
```bash
./network-config firewall list
./network-config firewall status
```

### Monitor Traffic Issues
```bash
./network-config monitor traffic --interface eth0
./network-config monitor stats
```

### QoS Verification
```bash
./network-config qos policy --name "test-policy" --bandwidth "100Mbps"
./network-config qos limit --interface eth0 --limit "50Mbps"
```

## Best Practices

1. **VLAN Design**
   - Isolate critical systems
   - Separate management traffic
   - Use guest VLAN for visitors

2. **VPN Security**
   - Use OpenVPN or L2TP for security-critical connections
   - Implement certificate-based authentication
   - Rotate VPN credentials regularly

3. **Firewall Rules**
   - Start with "deny all" policy
   - Explicitly allow necessary traffic
   - Log blocked connections
   - Review rules regularly

4. **QoS Configuration**
   - Prioritize critical applications
   - Reserve bandwidth for VoIP
   - Limit peer-to-peer traffic
   - Monitor and adjust regularly

5. **Traffic Monitoring**
   - Set up continuous monitoring
   - Configure alerts for anomalies
   - Maintain traffic logs
   - Analyze trends periodically
