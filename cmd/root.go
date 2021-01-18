package cmd

import (
	"os"

	"gocli/models"

	"github.com/spf13/cobra"
)

// RootCmd is root command
var RootCmd = &cobra.Command{
	Use:   "igmaker",
	Short: "gitignore maker by command line",
	Run: func(cmd *cobra.Command, args []string) {
		if err := writeBytes("/Users/kenminami/test.txt"); err != nil {
			os.Exit(1)
		}
	},
}

// Execute root command
func Execute() {
	RootCmd.Execute()
}

// writeBytes write bytes to target file
func writeBytes(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range models.Golang {
		b := []byte(line)
		_, err := file.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}
