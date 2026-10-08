package runner

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/YogiSunil/MakeUtility/internal/config"
)

func newTestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/missing", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	mux.HandleFunc("/slow", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(80 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	})
	return httptest.NewServer(mux)
}

func TestCheck(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	tests := []struct {
		name       string
		endpoint   config.Endpoint
		wantPassed bool
		wantStatus int
	}{
		{
			name:       "expected 200",
			endpoint:   config.Endpoint{URL: srv.URL + "/ok", Method: "GET", ExpectedStatus: 200},
			wantPassed: true,
			wantStatus: 200,
		},
		{
			name:       "expected 404",
			endpoint:   config.Endpoint{URL: srv.URL + "/missing", Method: "GET", ExpectedStatus: 404},
			wantPassed: true,
			wantStatus: 404,
		},
		{
			name:       "wrong status",
			endpoint:   config.Endpoint{URL: srv.URL + "/missing", Method: "GET", ExpectedStatus: 200},
			wantPassed: false,
			wantStatus: 404,
		},
		{
			name:       "too slow",
			endpoint:   config.Endpoint{URL: srv.URL + "/slow", Method: "GET", ExpectedStatus: 200, MaxLatencyMs: 20},
			wantPassed: false,
			wantStatus: 200,
		},
		{
			name:       "slow but allowed",
			endpoint:   config.Endpoint{URL: srv.URL + "/slow", Method: "GET", ExpectedStatus: 200, MaxLatencyMs: 2000},
			wantPassed: true,
			wantStatus: 200,
		},
		{
			name:       "server not reachable",
			endpoint:   config.Endpoint{URL: "http://127.0.0.1:1", Method: "GET", ExpectedStatus: 200},
			wantPassed: false,
			wantStatus: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Check(tt.endpoint, 2*time.Second)
			if got.Passed != tt.wantPassed {
				t.Errorf("Passed = %v, want %v (error: %q)", got.Passed, tt.wantPassed, got.Error)
			}
			if got.StatusCode != tt.wantStatus {
				t.Errorf("StatusCode = %d, want %d", got.StatusCode, tt.wantStatus)
			}
			if !tt.wantPassed && got.Error == "" {
				t.Error("a failed check should have an error message")
			}
		})
	}
}

func TestRunAllKeepsOrder(t *testing.T) {
	srv := newTestServer()
	defer srv.Close()

	endpoints := []config.Endpoint{
		{Name: "slow", URL: srv.URL + "/slow", Method: "GET", ExpectedStatus: 200},
		{Name: "ok", URL: srv.URL + "/ok", Method: "GET", ExpectedStatus: 200},
		{Name: "missing", URL: srv.URL + "/missing", Method: "GET", ExpectedStatus: 404},
	}

	results := RunAll(endpoints, 2*time.Second)
	if len(results) != len(endpoints) {
		t.Fatalf("got %d results, want %d", len(results), len(endpoints))
	}
	for i, e := range endpoints {
		if results[i].Name != e.Name {
			t.Errorf("result %d is %q, want %q", i, results[i].Name, e.Name)
		}
	}
}
