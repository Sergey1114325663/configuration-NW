package network

import (
	"fmt"
	"sync"
	"time"
)

// TrafficStats represents network traffic statistics
type TrafficStats struct {
	Interface   string
	RXBytes     uint64
	TXBytes     uint64
	RXPackets   uint64
	TXPackets   uint64
	RXErrors    uint64
	TXErrors    uint64
	Timestamp   time.Time
	BandwidthRX float64 // Mbps
	BandwidthTX float64 // Mbps
}

// TrafficMonitor monitors network traffic
type TrafficMonitor struct {
	stats map[string]*TrafficStats
	mu    sync.RWMutex
}

// NewTrafficMonitor creates a new traffic monitor
func NewTrafficMonitor() *TrafficMonitor {
	return &TrafficMonitor{
		stats: make(map[string]*TrafficStats),
	}
}

// UpdateStats updates traffic statistics
func (tm *TrafficMonitor) UpdateStats(iface string, stats *TrafficStats) {
	tm.mu.Lock()
	defer tm.mu.Unlock()
	stats.Timestamp = time.Now()
	tm.stats[iface] = stats
}

// GetStats retrieves traffic statistics for an interface
func (tm *TrafficMonitor) GetStats(iface string) (*TrafficStats, bool) {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	stats, ok := tm.stats[iface]
	return stats, ok
}

// GetAllStats retrieves all statistics
func (tm *TrafficMonitor) GetAllStats() map[string]*TrafficStats {
	tm.mu.RLock()
	defer tm.mu.RUnlock()
	// Create a copy to avoid race conditions
	result := make(map[string]*TrafficStats)
	for k, v := range tm.stats {
		result[k] = v
	}
	return result
}

// PrintStats prints traffic statistics
func (tm *TrafficMonitor) PrintStats(iface string) {
	stats, ok := tm.GetStats(iface)
	if !ok {
		fmt.Printf("No stats available for interface: %s\n", iface)
		return
	}

	fmt.Printf("\n=== Traffic Statistics for %s ===\n", iface)
	fmt.Printf("Time: %s\n", stats.Timestamp.Format("2006-01-02 15:04:05"))
	fmt.Printf("RX bytes: %d (%.2f MB)\n", stats.RXBytes, float64(stats.RXBytes)/1024/1024)
	fmt.Printf("TX bytes: %d (%.2f MB)\n", stats.TXBytes, float64(stats.TXBytes)/1024/1024)
	fmt.Printf("RX packets: %d\n", stats.RXPackets)
	fmt.Printf("TX packets: %d\n", stats.TXPackets)
	fmt.Printf("RX errors: %d\n", stats.RXErrors)
	fmt.Printf("TX errors: %d\n", stats.TXErrors)
	fmt.Printf("RX bandwidth: %.2f Mbps\n", stats.BandwidthRX)
	fmt.Printf("TX bandwidth: %.2f Mbps\n", stats.BandwidthTX)
}

// VLANInfo represents VLAN information
type VLANInfo struct {
	ID       string
	Name     string
	Status   string
	Members  []string
	MTU      int
	TaggedIF string
}

// FirewallStats represents firewall statistics
type FirewallStats struct {
	ActiveRules     int
	BlockedPackets  uint64
	AllowedPackets  uint64
	RejectedPackets uint64
	LastUpdated     time.Time
}

// QoSStatus represents QoS status
type QoSStatus struct {
	Interface   string
	Bandwidth   string
	Utilization float64
	DroppedPkts uint64
	Priority    int
}

// VPNConnection represents a VPN connection status
type VPNConnection struct {
	Name        string
	Type        string
	Status      string
	BytesIn     uint64
	BytesOut    uint64
	ConnectTime time.Time
	IPAddress   string
}

// ConnectionPool manages network connections
type ConnectionPool struct {
	connections map[string]*VPNConnection
	mu          sync.RWMutex
}

// NewConnectionPool creates a new connection pool
func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		connections: make(map[string]*VPNConnection),
	}
}

// AddConnection adds a connection to the pool
func (cp *ConnectionPool) AddConnection(conn *VPNConnection) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	cp.connections[conn.Name] = conn
}

// RemoveConnection removes a connection from the pool
func (cp *ConnectionPool) RemoveConnection(name string) {
	cp.mu.Lock()
	defer cp.mu.Unlock()
	delete(cp.connections, name)
}

// GetConnection retrieves a connection
func (cp *ConnectionPool) GetConnection(name string) (*VPNConnection, bool) {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	conn, ok := cp.connections[name]
	return conn, ok
}

// GetAllConnections retrieves all connections
func (cp *ConnectionPool) GetAllConnections() map[string]*VPNConnection {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	result := make(map[string]*VPNConnection)
	for k, v := range cp.connections {
		result[k] = v
	}
	return result
}

// ConnectionCount returns the number of active connections
func (cp *ConnectionPool) ConnectionCount() int {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return len(cp.connections)
}
