// Package cmd contains the APIForge command-line interface.
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "apiforge",
	Short:         "APIForge tests REST API endpoints",
	SilenceErrors: true,
	Long: `APIForge reads a list of REST API endpoints from a JSON file,
sends HTTP requests to each one, checks the status code and response
time, and saves a report.`,
}

// Execute runs the root command and exits with a non-zero code on error.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
