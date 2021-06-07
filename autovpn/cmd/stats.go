package cmd

import (
	"github.com/rogercoll/autovpn"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Statistics of current OpenVpn available connections",
	Long:  `This command ping all the available servers of your Openvpn configuration files and return the average round-trip time`,
	Run: func(cmd *cobra.Command, args []string) {
		configFile, err := cmd.Flags().GetString("vpns")
		if err != nil {
			panic(err)
		}
		getAutoVpn(configFile)
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
	statsCmd.PersistentFlags().StringP("vpns", "v", "/etc/autovpn/confs", "Openvpn credentials file (default is $HOME/.autovpn/openvpn.creds)")
}

func getAutoVpn(configFile string) {
	err := autovpn.Statistics(configFile)
	if err != nil {
		panic(err)
	}
}
