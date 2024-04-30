/*
Copyright © 2024 Mohammad Alsmadi <smadi.dev@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/smadixd/toolbox-cli/cmd/info"
	"github.com/smadixd/toolbox-cli/cmd/net"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "toolbox",
	Short: "A brief description of your application",
	Long:  ``,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func addSubcommandsPalettes() {
	rootCmd.AddCommand(net.NetCmd)
	rootCmd.AddCommand(info.InfoCmd)
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubcommandsPalettes()
}
