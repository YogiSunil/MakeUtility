package runner

import (
	"net/http"
	"time"

	"github.com/YogiSunil/MakeUtility/internal/config"
)

// Check sends one request to the endpoint and returns what happened.
func Check(e config.Endpoint) Result {
	result := Result{Name: e.Name, URL: e.URL, Method: e.Method}

	req, err := http.NewRequest(e.Method, e.URL, nil)
	if err != nil {
		result.Error = err.Error()
		return result
	}

	start := time.Now()
	resp, err := http.DefaultClient.Do(req)
	result.LatencyMs = time.Since(start).Milliseconds()
	if err != nil {
		result.Error = err.Error()
		return result
	}
	defer resp.Body.Close()

	result.StatusCode = resp.StatusCode
	return result
}
