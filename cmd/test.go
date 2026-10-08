package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test <endpoints.json>",
	Short: "Test the endpoints listed in a JSON file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("Testing endpoints from %s\n", args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
