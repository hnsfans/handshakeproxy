package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var proxyVersion = "0.0.1"

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version  of handshakeproxy",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Handshake Proxy v%s\n", proxyVersion)
	},
}
