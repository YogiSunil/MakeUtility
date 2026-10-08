package runner

// Result is the outcome of testing one endpoint.
type Result struct {
	Name       string `json:"name"`
	URL        string `json:"url"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	LatencyMs  int64  `json:"latency_ms"`
	Passed     bool   `json:"passed"`
	Error      string `json:"error,omitempty"`
}
