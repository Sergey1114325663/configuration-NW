package network

import (
	"crypto/tls"
	"fmt"
	"sync"
	"time"
)

// VPNManager manages VPN connections
type VPNManager struct {
	connections map[string]*VPNConnection
	mu          sync.RWMutex
}

// VPNConfig represents VPN configuration
type VPNConfig struct {
	Name          string
	Type          string // L2TP, PPTP, SSTP, OpenVPN
	Server        string
	Port          int
	Username      string
	Password      string
	EncryptionKey string
	Protocol      string
	MTU           int
	Compression   bool
	DualStack     bool
	DNS           []string
}

// NewVPNManager creates a new VPN manager
func NewVPNManager() *VPNManager {
	return &VPNManager{
		connections: make(map[string]*VPNConnection),
	}
}

// CreateConnection creates a new VPN connection
func (vm *VPNManager) CreateConnection(config VPNConfig) error {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	if _, exists := vm.connections[config.Name]; exists {
		return fmt.Errorf("connection already exists: %s", config.Name)
	}

	conn := &VPNConnection{
		Name:        config.Name,
		Type:        config.Type,
		Status:      "disconnected",
		ConnectTime: time.Now(),
	}

	vm.connections[config.Name] = conn
	return nil
}

// Connect establishes a VPN connection
func (vm *VPNManager) Connect(name string) error {
	vm.mu.Lock()
	conn, exists := vm.connections[name]
	vm.mu.Unlock()

	if !exists {
		return fmt.Errorf("connection not found: %s", name)
	}

	// Simulate connection establishment
	conn.Status = "connecting"
	time.Sleep(time.Second) // Simulate connection time

	vm.mu.Lock()
	conn.Status = "connected"
	conn.ConnectTime = time.Now()
	vm.mu.Unlock()

	return nil
}

// Disconnect closes a VPN connection
func (vm *VPNManager) Disconnect(name string) error {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	conn, exists := vm.connections[name]
	if !exists {
		return fmt.Errorf("connection not found: %s", name)
	}

	conn.Status = "disconnected"
	return nil
}

// GetConnection retrieves a VPN connection
func (vm *VPNManager) GetConnection(name string) (*VPNConnection, bool) {
	vm.mu.RLock()
	defer vm.mu.RUnlock()
	conn, ok := vm.connections[name]
	return conn, ok
}

// ListConnections returns all VPN connections
func (vm *VPNManager) ListConnections() []*VPNConnection {
	vm.mu.RLock()
	defer vm.mu.RUnlock()

	conns := make([]*VPNConnection, 0, len(vm.connections))
	for _, conn := range vm.connections {
		conns = append(conns, conn)
	}
	return conns
}

// DeleteConnection removes a VPN connection
func (vm *VPNManager) DeleteConnection(name string) error {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	if _, exists := vm.connections[name]; !exists {
		return fmt.Errorf("connection not found: %s", name)
	}

	delete(vm.connections, name)
	return nil
}

// L2TPConfig represents L2TP specific configuration
type L2TPConfig struct {
	Base               VPNConfig
	IPSecProfile       string // IKEv2, IKEv1
	EncryptionAlgo     string // AES-256, AES-128
	AuthenticationAlgo string // SHA-256, SHA-512
	PFSGroup           int    // 14, 15, 16
}

// PPTPConfig represents PPTP specific configuration
type PPTPConfig struct {
	Base              VPNConfig
	MPPEEncryption    bool
	EncryptionLevel   string // 40-bit, 56-bit, 128-bit
	Compression       bool
	StatelessCompress bool
}

// SSTPConfig represents SSTP specific configuration
type SSTPConfig struct {
	Base         VPNConfig
	SSLProtocol  string // TLS 1.2, TLS 1.3
	Certificate  string // Certificate path
	VerifyServer bool
	ProxySupport bool
}

// OpenVPNConfig represents OpenVPN specific configuration
type OpenVPNConfig struct {
	Base             VPNConfig
	ConfigFile       string
	CertificatePath  string
	KeyPath          string
	CAPath           string
	Cipher           string // AES-256-CBC
	Auth             string // SHA256
	Compression      string // LZO, Snappy
	KeyDirection     int
	ConnectTimeout   int
	ReconnectTimeout int
}

// VLANManager manages VLAN configuration
type VLANManager struct {
	vlans map[string]*VLANInfo
	mu    sync.RWMutex
}

// NewVLANManager creates a new VLAN manager
func NewVLANManager() *VLANManager {
	return &VLANManager{
		vlans: make(map[string]*VLANInfo),
	}
}

// CreateVLAN creates a new VLAN
func (vm *VLANManager) CreateVLAN(vlan *VLANInfo) error {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	if _, exists := vm.vlans[vlan.ID]; exists {
		return fmt.Errorf("VLAN already exists: %s", vlan.ID)
	}

	if vlan.MTU == 0 {
		vlan.MTU = 1500 // Default Ethernet MTU
	}

	vm.vlans[vlan.ID] = vlan
	return nil
}

// GetVLAN retrieves a VLAN
func (vm *VLANManager) GetVLAN(id string) (*VLANInfo, bool) {
	vm.mu.RLock()
	defer vm.mu.RUnlock()
	vlan, ok := vm.vlans[id]
	return vlan, ok
}

// ListVLANs returns all VLANs
func (vm *VLANManager) ListVLANs() []*VLANInfo {
	vm.mu.RLock()
	defer vm.mu.RUnlock()

	vlans := make([]*VLANInfo, 0, len(vm.vlans))
	for _, vlan := range vm.vlans {
		vlans = append(vlans, vlan)
	}
	return vlans
}

// DeleteVLAN removes a VLAN
func (vm *VLANManager) DeleteVLAN(id string) error {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	if _, exists := vm.vlans[id]; !exists {
		return fmt.Errorf("VLAN not found: %s", id)
	}

	delete(vm.vlans, id)
	return nil
}

// AddInterface adds an interface to a VLAN
func (vm *VLANManager) AddInterface(vlanID, ifName string) error {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	vlan, exists := vm.vlans[vlanID]
	if !exists {
		return fmt.Errorf("VLAN not found: %s", vlanID)
	}

	for _, member := range vlan.Members {
		if member == ifName {
			return fmt.Errorf("interface already in VLAN: %s", ifName)
		}
	}

	vlan.Members = append(vlan.Members, ifName)
	return nil
}

// RemoveInterface removes an interface from a VLAN
func (vm *VLANManager) RemoveInterface(vlanID, ifName string) error {
	vm.mu.Lock()
	defer vm.mu.Unlock()

	vlan, exists := vm.vlans[vlanID]
	if !exists {
		return fmt.Errorf("VLAN not found: %s", vlanID)
	}

	for i, member := range vlan.Members {
		if member == ifName {
			vlan.Members = append(vlan.Members[:i], vlan.Members[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("interface not found in VLAN: %s", ifName)
}

// SecurityConfig represents security settings
type SecurityConfig struct {
	TLSVersion         string // TLS1.2, TLS1.3
	CipherSuites       []string
	CertificatePath    string
	KeyPath            string
	CAPath             string
	VerifyPeer         bool
	VerifyHostname     bool
	InsecureSkipVerify bool
	Timeout            time.Duration
}

// GetTLSConfig returns TLS configuration
func (sc *SecurityConfig) GetTLSConfig() (*tls.Config, error) {
	tlsConfig := &tls.Config{
		InsecureSkipVerify: sc.InsecureSkipVerify,
	}

	// Set minimum version
	switch sc.TLSVersion {
	case "TLS1.2":
		tlsConfig.MinVersion = tls.VersionTLS12
	case "TLS1.3":
		tlsConfig.MinVersion = tls.VersionTLS13
	default:
		tlsConfig.MinVersion = tls.VersionTLS12
	}

	return tlsConfig, nil
}
