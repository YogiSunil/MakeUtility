package report

import (
	"fmt"
	"io"

	"github.com/YogiSunil/MakeUtility/internal/runner"
)

const divider = "------------------------------------------------"

// Print writes the results to w as a simple table.
func Print(w io.Writer, results []runner.Result, verbose bool) {
	fmt.Fprintln(w, "APIForge Results")
	fmt.Fprintln(w, divider)

	for _, r := range results {
		status := "PASS"
		if !r.Passed {
			status = "FAIL"
		}
		fmt.Fprintf(w, "%-6s %-24s %s %6dms\n", r.Method, r.Name, status, r.LatencyMs)
		if r.Error != "" {
			fmt.Fprintf(w, "       -> %s\n", r.Error)
		}
		if verbose {
			fmt.Fprintf(w, "       %s (status %d)\n", r.URL, r.StatusCode)
		}
	}

	fmt.Fprintln(w, divider)
	passed, failed := Summary(results)
	fmt.Fprintf(w, "Passed: %d | Failed: %d\n", passed, failed)
}
