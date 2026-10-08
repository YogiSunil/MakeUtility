package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/YogiSunil/MakeUtility/internal/config"
	"github.com/YogiSunil/MakeUtility/internal/report"
	"github.com/YogiSunil/MakeUtility/internal/runner"
)

var (
	timeout time.Duration
	output  string
)

var testCmd = &cobra.Command{
	Use:          "test <endpoints.json>",
	Short:        "Test the endpoints listed in a JSON file",
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load(args[0])
		if err != nil {
			return err
		}

		out := cmd.OutOrStdout()
		results := runner.RunAll(cfg.Endpoints, timeout)
		report.Print(out, results)

		if err := report.Save(output, results); err != nil {
			return err
		}
		fmt.Fprintf(out, "\nReport saved to %s\n", output)

		if _, failed := report.Summary(results); failed > 0 {
			return fmt.Errorf("%d endpoint(s) failed", failed)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.Flags().StringVarP(&output, "output", "o", "results.json", "where to save the JSON report")
	testCmd.Flags().DurationVar(&timeout, "timeout", runner.DefaultTimeout, "how long to wait for each endpoint")
}
