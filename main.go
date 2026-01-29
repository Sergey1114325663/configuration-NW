package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"network-config/cmd"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(logrus.InfoLevel)
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "network-config",
		Short: "Network Configuration Manager",
		Long:  "Manage VLAN, VPN, Firewall, QoS and traffic monitoring",
	}

	// Add subcommands
	rootCmd.AddCommand(cmd.VLANCmd())
	rootCmd.AddCommand(cmd.VPNCmd())
	rootCmd.AddCommand(cmd.FirewallCmd())
	rootCmd.AddCommand(cmd.QoSCmd())
	rootCmd.AddCommand(cmd.MonitorCmd())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
