package cmd

import "github.com/spf13/cobra"

var (
	rootCmd = &cobra.Command{Use: "handshakeproxy"}
	verbose bool
)

func init() {
	rootCmd.AddCommand(proxyCmd, debugCmd, versionCmd)
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")

}

// Main cli 模块入口函数
func Main() {
	rootCmd.Execute()
}
