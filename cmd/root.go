package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCmd is root command
var RootCmd = &cobra.Command{
	Use:   "igmaker",
	Short: "gitignore maker by command line",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command")
	},
}

// Execute root command
func Execute() {
	RootCmd.Execute()
}
