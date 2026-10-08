package config

import (
	"os"
	"path/filepath"
	"testing"
)

func writeTempFile(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "endpoints.json")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestLoad(t *testing.T) {
	tests := []struct {
		name      string
		content   string
		wantErr   bool
		wantCount int
	}{
		{
			name: "one endpoint",
			content: `{"endpoints": [
				{"name": "home", "url": "https://example.com", "method": "GET", "expected_status": 200, "max_latency_ms": 500}
			]}`,
			wantCount: 1,
		},
		{
			name: "two endpoints",
			content: `{"endpoints": [
				{"url": "https://example.com/a", "method": "GET"},
				{"url": "https://example.com/b", "method": "POST"}
			]}`,
			wantCount: 2,
		},
		{name: "broken json", content: `{"endpoints": [`, wantErr: true},
		{name: "empty list", content: `{"endpoints": []}`, wantErr: true},
		{name: "bad url", content: `{"endpoints": [{"url": "nope", "method": "GET"}]}`, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := Load(writeTempFile(t, tt.content))
			if (err != nil) != tt.wantErr {
				t.Fatalf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err == nil && len(cfg.Endpoints) != tt.wantCount {
				t.Errorf("got %d endpoints, want %d", len(cfg.Endpoints), tt.wantCount)
			}
		})
	}
}

func TestLoadMissingFile(t *testing.T) {
	if _, err := Load(filepath.Join(t.TempDir(), "nope.json")); err == nil {
		t.Error("expected an error for a missing file")
	}
}
