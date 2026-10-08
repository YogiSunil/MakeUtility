package runner

import (
	"fmt"
	"net/http"
	"time"

	"github.com/YogiSunil/MakeUtility/internal/config"
)

// DefaultTimeout is how long we wait for one endpoint to answer.
const DefaultTimeout = 10 * time.Second

// Check sends one request to the endpoint and returns what happened.
func Check(e config.Endpoint, timeout time.Duration) Result {
	result := Result{Name: e.Name, URL: e.URL, Method: e.Method}

	req, err := http.NewRequest(e.Method, e.URL, nil)
	if err != nil {
		result.Error = err.Error()
		return result
	}

	client := &http.Client{Timeout: timeout}
	start := time.Now()
	resp, err := client.Do(req)
	result.LatencyMs = time.Since(start).Milliseconds()
	if err != nil {
		result.Error = err.Error()
		return result
	}
	defer resp.Body.Close()

	result.StatusCode = resp.StatusCode
	switch {
	case resp.StatusCode != e.ExpectedStatus:
		result.Error = fmt.Sprintf("expected status %d, got %d", e.ExpectedStatus, resp.StatusCode)
	case e.MaxLatencyMs > 0 && result.LatencyMs > int64(e.MaxLatencyMs):
		result.Error = fmt.Sprintf("too slow: %dms (max %dms)", result.LatencyMs, e.MaxLatencyMs)
	default:
		result.Passed = true
	}
	return result
}
