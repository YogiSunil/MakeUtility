package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/YogiSunil/MakeUtility/internal/config"
	"github.com/YogiSunil/MakeUtility/internal/runner"
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

		results := runner.RunAll(cfg.Endpoints, runner.DefaultTimeout)
		for _, r := range results {
			fmt.Printf("%s %s %d %dms\n", r.Method, r.Name, r.StatusCode, r.LatencyMs)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
