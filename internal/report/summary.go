package report

import "github.com/YogiSunil/MakeUtility/internal/runner"

// Summary counts how many results passed and how many failed.
func Summary(results []runner.Result) (passed, failed int) {
	for _, r := range results {
		if r.Passed {
			passed++
		} else {
			failed++
		}
	}
	return passed, failed
}
