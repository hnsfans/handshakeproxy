package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{Use: "handshakeproxy"}

func init() {
	rootCmd.AddCommand(versionCmd, proxyCmd)
}

// Main cli 模块入口函数
func Main() {
	rootCmd.Execute()
}
