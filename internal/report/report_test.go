package report

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/YogiSunil/MakeUtility/internal/runner"
)

func sampleResults() []runner.Result {
	return []runner.Result{
		{Name: "users", Method: "GET", StatusCode: 200, LatencyMs: 120, Passed: true},
		{Name: "orders", Method: "GET", StatusCode: 500, LatencyMs: 350, Error: "expected status 200, got 500"},
		{Name: "products", Method: "GET", StatusCode: 200, LatencyMs: 85, Passed: true},
	}
}

func TestSummary(t *testing.T) {
	tests := []struct {
		name       string
		results    []runner.Result
		wantPassed int
		wantFailed int
	}{
		{"no results", nil, 0, 0},
		{"all passed", []runner.Result{{Passed: true}, {Passed: true}}, 2, 0},
		{"all failed", []runner.Result{{}, {}}, 0, 2},
		{"mixed", sampleResults(), 2, 1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			passed, failed := Summary(tt.results)
			if passed != tt.wantPassed || failed != tt.wantFailed {
				t.Errorf("Summary() = %d, %d; want %d, %d", passed, failed, tt.wantPassed, tt.wantFailed)
			}
		})
	}
}

func TestSave(t *testing.T) {
	path := filepath.Join(t.TempDir(), "results.json")
	if err := Save(path, sampleResults()); err != nil {
		t.Fatal(err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}

	var rep Report
	if err := json.Unmarshal(data, &rep); err != nil {
		t.Fatalf("saved file is not valid JSON: %v", err)
	}
	if rep.Passed != 2 || rep.Failed != 1 || len(rep.Results) != 3 {
		t.Errorf("unexpected report: passed=%d failed=%d results=%d", rep.Passed, rep.Failed, len(rep.Results))
	}
}

func TestPrint(t *testing.T) {
	var buf bytes.Buffer
	Print(&buf, sampleResults(), false)

	out := buf.String()
	for _, want := range []string{"PASS", "FAIL", "orders", "Passed: 2 | Failed: 1"} {
		if !strings.Contains(out, want) {
			t.Errorf("output is missing %q:\n%s", want, out)
		}
	}
}

func TestPrintVerbose(t *testing.T) {
	results := []runner.Result{
		{Name: "users", URL: "https://example.com/users", Method: "GET", StatusCode: 200, Passed: true},
	}

	var quiet, loud bytes.Buffer
	Print(&quiet, results, false)
	Print(&loud, results, true)

	if strings.Contains(quiet.String(), "https://example.com/users") {
		t.Error("the url should only be shown with verbose on")
	}
	if !strings.Contains(loud.String(), "https://example.com/users (status 200)") {
		t.Errorf("verbose output is missing the url and status:\n%s", loud.String())
	}
}
