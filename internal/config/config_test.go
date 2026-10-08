package config

import "testing"

func TestValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{
			name: "valid endpoint",
			cfg:  Config{Endpoints: []Endpoint{{Name: "ok", URL: "https://example.com", Method: "GET"}}},
		},
		{
			name: "lowercase method is accepted",
			cfg:  Config{Endpoints: []Endpoint{{URL: "https://example.com", Method: "get"}}},
		},
		{
			name:    "no endpoints",
			cfg:     Config{},
			wantErr: true,
		},
		{
			name:    "missing url",
			cfg:     Config{Endpoints: []Endpoint{{Name: "x", Method: "GET"}}},
			wantErr: true,
		},
		{
			name:    "not a url",
			cfg:     Config{Endpoints: []Endpoint{{URL: "not a url", Method: "GET"}}},
			wantErr: true,
		},
		{
			name:    "unsupported scheme",
			cfg:     Config{Endpoints: []Endpoint{{URL: "ftp://example.com", Method: "GET"}}},
			wantErr: true,
		},
		{
			name:    "unsupported method",
			cfg:     Config{Endpoints: []Endpoint{{URL: "https://example.com", Method: "FETCH"}}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
