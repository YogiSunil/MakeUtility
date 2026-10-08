package config

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Load reads the endpoints file at path.
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("reading config: %w", err)
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("parsing %s: %w", path, err)
	}

	cfg.applyDefaults()

	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return &cfg, nil
}

// applyDefaults fills in values the user left out of the file.
func (c *Config) applyDefaults() {
	for i := range c.Endpoints {
		e := &c.Endpoints[i]
		if e.Method == "" {
			e.Method = http.MethodGet
		}
		e.Method = strings.ToUpper(e.Method)
		if e.ExpectedStatus == 0 {
			e.ExpectedStatus = http.StatusOK
		}
		if e.Name == "" {
			e.Name = e.URL
		}
	}
}
