package runner

import (
	"time"

	"github.com/YogiSunil/MakeUtility/internal/config"
)

// RunAll checks the endpoints one after another.
func RunAll(endpoints []config.Endpoint, timeout time.Duration) []Result {
	results := make([]Result, 0, len(endpoints))
	for _, e := range endpoints {
		results = append(results, Check(e, timeout))
	}
	return results
}
