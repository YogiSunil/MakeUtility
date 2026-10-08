package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/YogiSunil/MakeUtility/internal/config"
)

var testCmd = &cobra.Command{
	Use:   "test <endpoints.json>",
	Short: "Test the endpoints listed in a JSON file",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load(args[0])
		if err != nil {
			return err
		}

		fmt.Printf("Loaded %d endpoints from %s\n", len(cfg.Endpoints), args[0])
		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
