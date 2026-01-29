package tests

import (
	"fmt"
	"network-config/network"
	"testing"
)

// TestVLANManager tests VLAN functionality
func TestVLANManager(t *testing.T) {
	vm := network.NewVLANManager()

	// Test Create VLAN
	vlan := &network.VLANInfo{
		ID:       "10",
		Name:     "Production",
		Status:   "active",
		Members:  []string{"eth0"},
		MTU:      1500,
		TaggedIF: "eth0.10",
	}

	if err := vm.CreateVLAN(vlan); err != nil {
		t.Fatalf("Failed to create VLAN: %v", err)
	}

	// Test Get VLAN
	retrieved, ok := vm.GetVLAN("10")
	if !ok {
		t.Fatal("VLAN not found")
	}

	if retrieved.Name != "Production" {
		t.Fatalf("Expected name 'Production', got '%s'", retrieved.Name)
	}

	// Test List VLANs
	vlans := vm.ListVLANs()
	if len(vlans) != 1 {
		t.Fatalf("Expected 1 VLAN, got %d", len(vlans))
	}

	// Test Add Interface
	if err := vm.AddInterface("10", "eth1"); err != nil {
		t.Fatalf("Failed to add interface: %v", err)
	}

	// Test Delete VLAN
	if err := vm.DeleteVLAN("10"); err != nil {
		t.Fatalf("Failed to delete VLAN: %v", err)
	}

	fmt.Println("✓ VLAN Manager tests passed")
}

// TestFirewallManager tests Firewall functionality
func TestFirewallManager(t *testing.T) {
	fm := network.NewFirewallManager()

	// Test IsEnabled
	if !fm.IsEnabled() {
		t.Fatal("Firewall should be enabled")
	}

	// Test Disable
	fm.Disable()
	if fm.IsEnabled() {
		t.Fatal("Firewall should be disabled")
	}

	// Test Enable
	fm.Enable()
	if !fm.IsEnabled() {
		t.Fatal("Firewall should be enabled")
	}

	fmt.Println("✓ Firewall Manager tests passed")
}

// TestQoSManager tests QoS functionality
func TestQoSManager(t *testing.T) {
	qm := network.NewQoSManager()

	// Test Create Policy
	policy := network.QoSPolicy{
		Name:      "high-priority",
		Interface: "eth0",
		Bandwidth: 100000, // 100Mbps
		Priority:  7,
	}

	if err := qm.CreatePolicy(policy); err != nil {
		t.Fatalf("Failed to create policy: %v", err)
	}

	// Test Get Policy
	retrieved, ok := qm.GetPolicy("high-priority")
	if !ok {
		t.Fatal("Policy not found")
	}

	if retrieved.Priority != 7 {
		t.Fatalf("Expected priority 7, got %d", retrieved.Priority)
	}

	// Test List Policies
	policies := qm.ListPolicies()
	if len(policies) != 1 {
		t.Fatalf("Expected 1 policy, got %d", len(policies))
	}

	// Test Delete Policy
	if err := qm.DeletePolicy("high-priority"); err != nil {
		t.Fatalf("Failed to delete policy: %v", err)
	}

	fmt.Println("✓ QoS Manager tests passed")
}

// TestVPNManager tests VPN functionality
func TestVPNManager(t *testing.T) {
	vm := network.NewVPNManager()

	// Test Create Connection
	config := network.VPNConfig{
		Name:   "office-vpn",
		Type:   "L2TP",
		Server: "vpn.example.com",
		Port:   1194,
	}

	if err := vm.CreateConnection(config); err != nil {
		t.Fatalf("Failed to create connection: %v", err)
	}

	// Test Get Connection
	conn, ok := vm.GetConnection("office-vpn")
	if !ok {
		t.Fatal("Connection not found")
	}

	if conn.Type != "L2TP" {
		t.Fatalf("Expected type 'L2TP', got '%s'", conn.Type)
	}

	fmt.Println("✓ VPN Manager tests passed")
}

// TestTrafficMonitor tests Traffic Monitoring
func TestTrafficMonitor(t *testing.T) {
	tm := network.NewTrafficMonitor()

	// Test Update Stats
	stats := &network.TrafficStats{
		Interface:   "eth0",
		RXBytes:     1000000,
		TXBytes:     500000,
		RXPackets:   10000,
		TXPackets:   5000,
		RXErrors:    0,
		TXErrors:    0,
		BandwidthRX: 80.5,
		BandwidthTX: 40.2,
	}

	tm.UpdateStats("eth0", stats)

	// Test Get Stats
	retrieved, ok := tm.GetStats("eth0")
	if !ok {
		t.Fatal("Stats not found")
	}

	if retrieved.RXBytes != 1000000 {
		t.Fatalf("Expected RX bytes 1000000, got %d", retrieved.RXBytes)
	}

	fmt.Println("✓ Traffic Monitor tests passed")
}

// TestConnectionPool tests Connection Pool
func TestConnectionPool(t *testing.T) {
	cp := network.NewConnectionPool()

	// Test Add Connection
	conn := &network.VPNConnection{
		Name:   "test-vpn",
		Type:   "OpenVPN",
		Status: "connected",
	}

	cp.AddConnection(conn)

	// Test Get Connection
	retrieved, ok := cp.GetConnection("test-vpn")
	if !ok {
		t.Fatal("Connection not found")
	}

	if retrieved.Status != "connected" {
		t.Fatalf("Expected status 'connected', got '%s'", retrieved.Status)
	}

	// Test Connection Count
	if cp.ConnectionCount() != 1 {
		t.Fatalf("Expected 1 connection, got %d", cp.ConnectionCount())
	}

	// Test Remove Connection
	cp.RemoveConnection("test-vpn")

	if cp.ConnectionCount() != 0 {
		t.Fatalf("Expected 0 connections, got %d", cp.ConnectionCount())
	}

	fmt.Println("✓ Connection Pool tests passed")
}

// Run all tests
func TestAll(t *testing.T) {
	t.Run("VLAN Manager", TestVLANManager)
	t.Run("Firewall Manager", TestFirewallManager)
	t.Run("QoS Manager", TestQoSManager)
	t.Run("VPN Manager", TestVPNManager)
	t.Run("Traffic Monitor", TestTrafficMonitor)
	t.Run("Connection Pool", TestConnectionPool)
}
