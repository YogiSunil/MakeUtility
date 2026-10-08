package report

import (
	"testing"

	"github.com/YogiSunil/MakeUtility/internal/runner"
)

func BenchmarkSummary(b *testing.B) {
	results := make([]runner.Result, 1000)
	for i := range results {
		results[i] = runner.Result{Name: "endpoint", Passed: i%3 != 0}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Summary(results)
	}
}
