package config

import (
	"os"

	yaml "gopkg.in/yaml.v3"
)

type VLANConfig struct {
	ID          string   `yaml:"id"`
	Name        string   `yaml:"name"`
	Description string   `yaml:"description"`
	Interfaces  []string `yaml:"interfaces"`
	Status      string   `yaml:"status"`
}

type VPNConfig struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"` // L2TP, PPTP, SSTP, OpenVPN
	Server   string `yaml:"server"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Enabled  bool   `yaml:"enabled"`
}

type FirewallRule struct {
	ID        string `yaml:"id"`
	Direction string `yaml:"direction"` // Inbound, Outbound
	Action    string `yaml:"action"`    // Allow, Deny, Reject
	Protocol  string `yaml:"protocol"`  // TCP, UDP, ICMP
	Port      string `yaml:"port"`
	Source    string `yaml:"source"`
	Dest      string `yaml:"dest"`
}

type QoSPolicy struct {
	Name      string `yaml:"name"`
	Interface string `yaml:"interface"`
	Bandwidth string `yaml:"bandwidth"`
	Priority  int    `yaml:"priority"`
}

type NetworkConfig struct {
	VLANs      []VLANConfig   `yaml:"vlans"`
	VPN        []VPNConfig    `yaml:"vpn"`
	Firewall   []FirewallRule `yaml:"firewall"`
	QoS        []QoSPolicy    `yaml:"qos"`
	Monitoring struct {
		Enabled  bool   `yaml:"enabled"`
		Interval int    `yaml:"interval"`
		LogFile  string `yaml:"logfile"`
	} `yaml:"monitoring"`
}

// LoadConfig loads configuration from YAML file
func LoadConfig(filePath string) (*NetworkConfig, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config NetworkConfig
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// SaveConfig saves configuration to YAML file
func SaveConfig(config *NetworkConfig, filePath string) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
