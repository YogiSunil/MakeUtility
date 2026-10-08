package report

import (
	"fmt"
	"io"
	"strings"

	"github.com/YogiSunil/MakeUtility/internal/runner"
)

const divider = "------------------------------------------------"

// Print writes the results to w as a simple table.
func Print(w io.Writer, results []runner.Result, verbose bool) error {
	lines := []string{"APIForge Results", divider}

	for _, r := range results {
		status := "PASS"
		if !r.Passed {
			status = "FAIL"
		}
		lines = append(lines, fmt.Sprintf("%-6s %-24s %s %6dms", r.Method, r.Name, status, r.LatencyMs))
		if r.Error != "" {
			lines = append(lines, "       -> "+r.Error)
		}
		if verbose {
			lines = append(lines, fmt.Sprintf("       %s (status %d)", r.URL, r.StatusCode))
		}
	}

	lines = append(lines, divider)
	passed, failed := Summary(results)
	lines = append(lines, fmt.Sprintf("Passed: %d | Failed: %d", passed, failed))
	_, err := io.WriteString(w, strings.Join(lines, "\n")+"\n")
	return err
}
