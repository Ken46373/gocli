package cmd

import (
	"errors"
	"os"

	"gocli/models"

	"github.com/spf13/cobra"
)

// RootCmd is root command
var RootCmd = &cobra.Command{
	Use:   "igmaker",
	Short: "gitignore maker by command line",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("requires language name used in your project")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := writeBytes("/Users/kenminami/test.txt", args); err != nil {
			os.Exit(1)
		}
		return nil
	},
}

// Execute root command
func Execute() {
	RootCmd.Execute()
}

// writeBytes write bytes to target file
func writeBytes(filename string, args []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var model = findModel(args)
	for _, line := range model {
		b := []byte(line)
		_, err := file.Write(b)
		if err != nil {
			return err
		}
	}
	return nil
}

// findModel find language file
func findModel(args []string) []string {

	// todo
	return models.Golang
}
