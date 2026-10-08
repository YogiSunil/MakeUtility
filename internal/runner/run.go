package runner

import (
	"sync"
	"time"

	"github.com/YogiSunil/MakeUtility/internal/config"
)

// RunAll checks all endpoints at the same time. The results come back in
// the same order as the endpoints.
func RunAll(endpoints []config.Endpoint, timeout time.Duration) []Result {
	results := make([]Result, len(endpoints))

	var wg sync.WaitGroup
	for i, e := range endpoints {
		wg.Add(1)
		go func(i int, e config.Endpoint) {
			defer wg.Done()
			results[i] = Check(e, timeout)
		}(i, e)
	}
	wg.Wait()

	return results
}
