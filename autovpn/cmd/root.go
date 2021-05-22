package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "autovpn",
	Short: "An OpenVpn manager",
	Long:  `OpenVpn will manage your vpn configurations and will auto start.`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
