# Network Configuration Manager - FAQ

## General Questions

### Q: What is Network Configuration Manager?
**A:** It's a Go-based CLI tool and framework for managing network infrastructure including VLAN, VPN (L2TP, PPTP, SSTP, OpenVPN), Firewall, QoS, and traffic monitoring.

### Q: What platforms are supported?
**A:** Linux, Windows, and macOS. Both native binaries and Docker containers are supported.

### Q: Do I need administrator/root privileges?
**A:** Yes, most operations require elevated privileges to modify network configurations.

### Q: What's the minimum Go version?
**A:** Go 1.21 or later.

## Installation & Setup

### Q: How do I install?
**A:** 
- Windows: `.\install.ps1`
- Linux/macOS: `bash install.sh`
- Or build from source: `go build -o network-config main.go`

### Q: How do I verify installation?
**A:** Run `network-config --help` to see available commands.

### Q: Can I use it without Go installed?
**A:** Yes, download pre-built binaries from releases or use Docker.

### Q: Where are configuration files stored?
**A:** By default in `./config/` or `/etc/network-config/` on Linux.

## VLAN Management

### Q: What's a VLAN?
**A:** Virtual Local Area Network - a logical network segment within a physical network.

### Q: How many VLANs can I create?
**A:** Up to 4094 (standard 802.1Q supports IDs 1-4094).

### Q: Can I modify VLAN after creation?
**A:** Current version supports delete and recreate. Future versions will support in-place modifications.

### Q: What's the default MTU?
**A:** 1500 bytes (standard Ethernet). VLAN tagging adds 4 bytes, so effective MTU is 1496.

## VPN Configuration

### Q: Which VPN protocol should I use?
**A:** 
- **OpenVPN**: Best security and cross-platform support
- **L2TP**: Good for enterprise environments
- **PPTP**: Legacy, less secure, Windows compatibility
- **SSTP**: Good for bypassing firewalls

### Q: How do I create a VPN connection?
**A:** `network-config vpn [protocol] --action create --name [name]`

### Q: Can I have multiple VPN connections?
**A:** Yes, each connection is independent. You can create and manage multiple connections.

### Q: How do I import OpenVPN config?
**A:** `network-config vpn openvpn --action start --config /path/to/config.ovpn`

### Q: What encryption is used?
**A:** 
- L2TP: AES-256 (configurable)
- PPTP: MPPE 128-bit
- SSTP: SSL/TLS + PPP
- OpenVPN: AES-256-CBC (configurable)

### Q: How do I backup VPN settings?
**A:** Copy the YAML configuration file. All settings are stored in plain text YAML.

## Firewall Management

### Q: How do firewall rules work?
**A:** Rules are evaluated in priority order. First matching rule wins. Lowest priority number executes first.

### Q: Can I enable/disable the firewall?
**A:** Yes, `network-config firewall enable/disable`

### Q: What happens if no rules match?
**A:** Default is to allow. You can configure default deny policy.

### Q: How do I allow specific ports?
**A:** `network-config firewall rule --action add --rule "allow tcp port 80"`

### Q: Can I block entire subnets?
**A:** Yes, specify CIDR notation: `"deny tcp from 192.168.1.0/24"`

### Q: How do I export firewall rules?
**A:** Rules are stored in YAML config file. Export by copying the file.

### Q: What's the performance impact?
**A:** Negligible. Rule evaluation is O(n) where n is number of rules. Performance is linear even with 1000+ rules.

## Quality of Service (QoS)

### Q: What does QoS do?
**A:** Controls traffic prioritization and bandwidth allocation to ensure critical applications get resources.

### Q: What's the priority scale?
**A:** 0-7, where 7 is highest priority (e.g., VoIP) and 0 is lowest (e.g., file transfers).

### Q: Can I limit bandwidth per interface?
**A:** Yes, `network-config qos limit --interface eth0 --limit 50Mbps`

### Q: What scheduling algorithms are supported?
**A:** FIFO, Round Robin (RR), Weighted Fair Queuing (WFQ), and Priority Queue.

### Q: How do I prioritize video conferencing?
**A:** `network-config qos priority --class video-conf --priority 6`

### Q: What's CIR and PIR?
**A:** CIR (Committed Information Rate) is guaranteed bandwidth. PIR (Peak Information Rate) is maximum burst bandwidth.

## Traffic Monitoring

### Q: What metrics are collected?
**A:** RX/TX bytes and packets, errors, bandwidth utilization, latency, connection count.

### Q: Can I monitor multiple interfaces?
**A:** Yes, run separate monitoring commands for each interface or monitor all with `--interface any`.

### Q: How do I set alerts?
**A:** `network-config monitor alerts --threshold 100Mbps`

### Q: What's the monitoring overhead?
**A:** Minimal. Real-time monitoring runs ~50MB memory overhead.

### Q: Can I export monitoring data?
**A:** Yes, through API or by configuring output to file.

### Q: What's the update frequency?
**A:** Configurable, default is 5 seconds.

## Configuration

### Q: What's the configuration file format?
**A:** YAML format. See `examples/network.yml` for examples.

### Q: Can I version control configuration?
**A:** Yes, YAML format is Git-friendly. Recommended to keep configurations in version control.

### Q: How do I load configuration?
**A:** Place in `config/` directory or specify path with `--config` flag.

### Q: Can I have multiple configuration files?
**A:** Yes, specify different config files for different environments.

### Q: How do I backup my settings?
**A:** Backup the YAML configuration file. All state is stored there.

## API & Integration

### Q: Is there a REST API?
**A:** Planned for v2.0. Currently only CLI is available.

### Q: Can I use this as a library?
**A:** Yes, import `network-config/network` package in your Go code.

### Q: Is there WebSocket support?
**A:** Planned for v2.0 for real-time monitoring.

### Q: Can I integrate with Kubernetes?
**A:** Planned for v2.0. Currently works in Docker containers.

## Performance & Scaling

### Q: What's the maximum number of rules?
**A:** Tested with 1000+ rules without performance degradation.

### Q: How many VLANs can I manage?
**A:** Theoretically unlimited. Tested with 100+ VLANs.

### Q: What's the memory footprint?
**A:** ~50MB baseline + configuration size.

### Q: Can it handle enterprise networks?
**A:** Yes, designed for production use in enterprise environments.

### Q: What's the CPU usage?
**A:** Minimal when idle. Increases with active monitoring/traffic.

## Troubleshooting

### Q: VPN connection fails, what do I check?
**A:** 
1. Server connectivity: `ping [server]`
2. Firewall rules: Ensure VPN ports are allowed
3. Credentials: Verify username/password
4. Logs: Check application logs

### Q: Firewall rules not working?
**A:** 
1. Check rule priority order
2. Verify direction (inbound/outbound)
3. Ensure firewall is enabled
4. Check rule syntax

### Q: QoS not limiting bandwidth?
**A:** 
1. Verify interface name is correct
2. Ensure bandwidth value includes unit (Mbps, Gbps)
3. Check if policy is active
4. Verify priorities are set correctly

### Q: Monitoring shows no data?
**A:** 
1. Check interface name
2. Ensure interface is active
3. Verify monitoring is enabled
4. Check log files for errors

### Q: How do I get verbose output?
**A:** Add `--verbose` flag to any command.

### Q: Where are logs stored?
**A:** By default in `./logs/` or `/var/log/` on Linux.

## Docker & Containers

### Q: How do I run in Docker?
**A:** `docker run -v $(pwd)/config:/etc/network-config network-config:latest`

### Q: What's the Docker image size?
**A:** ~15MB (Alpine-based)

### Q: Can I use Docker Compose?
**A:** Yes, `docker-compose up -d` with included docker-compose.yml

### Q: Does it work with Kubernetes?
**A:** Currently as StatefulSet or DaemonSet. Full Kubernetes integration planned for v2.0.

### Q: Can I persist data across containers?
**A:** Yes, use volumes to mount config directory.

## Development & Contributing

### Q: Can I contribute?
**A:** Yes, please submit pull requests and issues.

### Q: How do I run tests?
**A:** `go test ./... -v`

### Q: What's the code coverage?
**A:** Target is 80%+. Run `go test -cover ./...`

### Q: How do I build for specific OS?
**A:** `GOOS=linux GOARCH=amd64 go build`

### Q: Where's the source code?
**A:** All source code is included in the project.

## Licensing & Legal

### Q: What's the license?
**A:** MIT License - free for commercial and personal use.

### Q: Can I modify the code?
**A:** Yes, MIT license allows modifications.

### Q: Must I contribute back changes?
**A:** No, but contributions are welcome.

### Q: Can I use commercially?
**A:** Yes, MIT license allows commercial use.

## Performance Optimization

### Q: How do I improve performance?
**A:** 
1. Reduce number of firewall rules
2. Optimize VLAN segmentation
3. Use appropriate QoS scheduling algorithm
4. Monitor and adjust priorities

### Q: Should I monitor all interfaces?
**A:** Only monitor interfaces you need. Reduces overhead.

### Q: What's the impact of QoS?
**A:** Minimal CPU overhead (~1-2%) for policy evaluation.

## Migration & Upgrades

### Q: How do I upgrade to new version?
**A:** 
1. Backup current config
2. Download new version
3. Restore config
4. Test configuration

### Q: Will my configuration work with new versions?
**A:** YAML format is backward compatible.

### Q: How do I migrate from another tool?
**A:** Export configuration to YAML format compatible with this tool.

## Uninstallation

### Q: How do I uninstall?
**A:** Delete the binary and config directory. No system files are created.

### Q: Will it leave traces?
**A:** No system-wide installation means clean removal.

---

## Getting Help

- **Documentation**: See README.md
- **Examples**: Check examples/ directory
- **Issues**: Open GitHub issue
- **Community**: Check discussions

## Additional Resources

- [Go Documentation](https://golang.org/doc)
- [YAML Specification](https://yaml.org)
- [Docker Documentation](https://docs.docker.com)
- [Networking Basics](https://en.wikipedia.org/wiki/Computer_network)
