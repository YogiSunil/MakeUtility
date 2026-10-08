package report

import (
	"encoding/json"
	"os"
	"time"

	"github.com/YogiSunil/MakeUtility/internal/runner"
)

// Report is what gets written to the JSON file.
type Report struct {
	GeneratedAt time.Time       `json:"generated_at"`
	Passed      int             `json:"passed"`
	Failed      int             `json:"failed"`
	Results     []runner.Result `json:"results"`
}

// Save writes the results and their totals to a JSON file.
func Save(path string, results []runner.Result) error {
	passed, failed := Summary(results)
	rep := Report{
		GeneratedAt: time.Now(),
		Passed:      passed,
		Failed:      failed,
		Results:     results,
	}

	data, err := json.MarshalIndent(rep, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}
