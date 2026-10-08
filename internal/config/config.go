// Package config defines the endpoint configuration read from JSON.
package config

// Endpoint describes one API endpoint to test.
type Endpoint struct {
	Name           string `json:"name"`
	URL            string `json:"url"`
	Method         string `json:"method"`
	ExpectedStatus int    `json:"expected_status"`
	MaxLatencyMs   int    `json:"max_latency_ms"`
}

// Config is the top-level structure of an endpoints file.
type Config struct {
	Endpoints []Endpoint `json:"endpoints"`
}
