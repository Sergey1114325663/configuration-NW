package network

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// FirewallManager manages firewall rules
type FirewallManager struct {
	rules   []FirewallRule
	enabled bool
	mu      sync.RWMutex
	stats   *FirewallStats
}

// FirewallRule represents a firewall rule
type FirewallRule struct {
	ID        string
	Direction string // "inbound" or "outbound"
	Action    string // "allow", "deny", "reject"
	Protocol  string // "tcp", "udp", "icmp"
	Port      string
	Source    net.IPNet
	Dest      net.IPNet
	Priority  int
}

// NewFirewallManager creates a new firewall manager
func NewFirewallManager() *FirewallManager {
	return &FirewallManager{
		rules:   make([]FirewallRule, 0),
		enabled: true,
		stats: &FirewallStats{
			LastUpdated: time.Now(),
		},
	}
}

// AddRule adds a firewall rule
func (fm *FirewallManager) AddRule(rule FirewallRule) error {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	// Validate rule
	if err := validateRule(rule); err != nil {
		return err
	}

	fm.rules = append(fm.rules, rule)
	fm.sortRulesByPriority()
	return nil
}

// RemoveRule removes a firewall rule by ID
func (fm *FirewallManager) RemoveRule(id string) error {
	fm.mu.Lock()
	defer fm.mu.Unlock()

	for i, rule := range fm.rules {
		if rule.ID == id {
			fm.rules = append(fm.rules[:i], fm.rules[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("rule not found: %s", id)
}

// GetRule retrieves a rule by ID
func (fm *FirewallManager) GetRule(id string) (FirewallRule, error) {
	fm.mu.RLock()
	defer fm.mu.RUnlock()

	for _, rule := range fm.rules {
		if rule.ID == id {
			return rule, nil
		}
	}
	return FirewallRule{}, fmt.Errorf("rule not found: %s", id)
}

// ListRules returns all firewall rules
func (fm *FirewallManager) ListRules() []FirewallRule {
	fm.mu.RLock()
	defer fm.mu.RUnlock()
	return fm.rules
}

// Enable enables the firewall
func (fm *FirewallManager) Enable() {
	fm.mu.Lock()
	defer fm.mu.Unlock()
	fm.enabled = true
}

// Disable disables the firewall
func (fm *FirewallManager) Disable() {
	fm.mu.Lock()
	defer fm.mu.Unlock()
	fm.enabled = false
}

// IsEnabled checks if firewall is enabled
func (fm *FirewallManager) IsEnabled() bool {
	fm.mu.RLock()
	defer fm.mu.RUnlock()
	return fm.enabled
}

// GetStats returns firewall statistics
func (fm *FirewallManager) GetStats() *FirewallStats {
	fm.mu.RLock()
	defer fm.mu.RUnlock()
	return fm.stats
}

// sortRulesByPriority sorts rules by priority
func (fm *FirewallManager) sortRulesByPriority() {
	// Simple bubble sort for demonstration
	for i := 0; i < len(fm.rules)-1; i++ {
		for j := 0; j < len(fm.rules)-i-1; j++ {
			if fm.rules[j].Priority > fm.rules[j+1].Priority {
				fm.rules[j], fm.rules[j+1] = fm.rules[j+1], fm.rules[j]
			}
		}
	}
}

// validateRule validates a firewall rule
func validateRule(rule FirewallRule) error {
	if rule.ID == "" {
		return fmt.Errorf("rule ID cannot be empty")
	}

	if rule.Direction != "inbound" && rule.Direction != "outbound" {
		return fmt.Errorf("invalid direction: %s", rule.Direction)
	}

	if rule.Action != "allow" && rule.Action != "deny" && rule.Action != "reject" {
		return fmt.Errorf("invalid action: %s", rule.Action)
	}

	if rule.Protocol != "tcp" && rule.Protocol != "udp" && rule.Protocol != "icmp" {
		return fmt.Errorf("invalid protocol: %s", rule.Protocol)
	}

	return nil
}

// QoSManager manages Quality of Service
type QoSManager struct {
	policies map[string]QoSPolicy
	mu       sync.RWMutex
}

// QoSPolicy represents a QoS policy
type QoSPolicy struct {
	Name        string
	Interface   string
	Bandwidth   uint64 // in Kbps
	Priority    int
	Algorithm   string // "FIFO", "RR", "WFQ"
	CIR         uint64 // Committed Information Rate
	PIR         uint64 // Peak Information Rate
	BurstSize   uint64
	DroppedPkts uint64
}

// NewQoSManager creates a new QoS manager
func NewQoSManager() *QoSManager {
	return &QoSManager{
		policies: make(map[string]QoSPolicy),
	}
}

// CreatePolicy creates a new QoS policy
func (qm *QoSManager) CreatePolicy(policy QoSPolicy) error {
	qm.mu.Lock()
	defer qm.mu.Unlock()

	if _, exists := qm.policies[policy.Name]; exists {
		return fmt.Errorf("policy already exists: %s", policy.Name)
	}

	qm.policies[policy.Name] = policy
	return nil
}

// DeletePolicy deletes a QoS policy
func (qm *QoSManager) DeletePolicy(name string) error {
	qm.mu.Lock()
	defer qm.mu.Unlock()

	if _, exists := qm.policies[name]; !exists {
		return fmt.Errorf("policy not found: %s", name)
	}

	delete(qm.policies, name)
	return nil
}

// GetPolicy retrieves a QoS policy
func (qm *QoSManager) GetPolicy(name string) (QoSPolicy, bool) {
	qm.mu.RLock()
	defer qm.mu.RUnlock()
	policy, ok := qm.policies[name]
	return policy, ok
}

// ListPolicies returns all QoS policies
func (qm *QoSManager) ListPolicies() map[string]QoSPolicy {
	qm.mu.RLock()
	defer qm.mu.RUnlock()
	return qm.policies
}

// UpdatePolicy updates a QoS policy
func (qm *QoSManager) UpdatePolicy(policy QoSPolicy) error {
	qm.mu.Lock()
	defer qm.mu.Unlock()

	if _, exists := qm.policies[policy.Name]; !exists {
		return fmt.Errorf("policy not found: %s", policy.Name)
	}

	qm.policies[policy.Name] = policy
	return nil
}
