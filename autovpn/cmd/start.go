package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/rogercoll/autovpn"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start new random OpenVpn connection",
	Long:  `This command fetches a random dad joke from the icanhazdadjoke api`,
	Run: func(cmd *cobra.Command, args []string) {
		credsFile, err := cmd.Flags().GetString("creds")
		if err != nil {
			panic(err)
		}
		configFile, err := cmd.Flags().GetString("vpns")
		if err != nil {
			panic(err)
		}
		startAutoVpn(credsFile, configFile)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	home, _ := homedir.Dir()
	startCmd.PersistentFlags().StringP("creds", "c", home+"/.autovpn/openvpn.creds", "Openvpn credentials file (default is $HOME/.autovpn/openvpn.creds)")
	startCmd.PersistentFlags().StringP("vpns", "v", home+"/.autovpn/confs", "Openvpn credentials file (default is $HOME/.autovpn/openvpn.creds)")
}

// StdoutLogger represents the stdout logger callback
type StdoutLogger func(text string)

// Log logs the given string to stdout logger
func (lc StdoutLogger) Log(text string) {
	lc(text)
}

func startAutoVpn(credsFile, configFile string) {
	var logger StdoutLogger = func(text string) {
		lines := strings.Split(text, "\n")
		for _, line := range lines {
			fmt.Println("Library check >>", line)
		}
	}
	file, err := os.Open(credsFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	username := scanner.Text()
	scanner.Scan()
	password := scanner.Text()
	conf, err := autovpn.NewConfig(username, password, configFile, logger, nil)
	if err != nil {
		panic(err)
	}
	conf.Start()
}
