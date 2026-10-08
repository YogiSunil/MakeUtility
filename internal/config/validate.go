package config

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

var validMethods = map[string]bool{
	"GET":    true,
	"POST":   true,
	"PUT":    true,
	"PATCH":  true,
	"DELETE": true,
	"HEAD":   true,
}

// Validate checks that every endpoint can actually be tested.
func (c *Config) Validate() error {
	if len(c.Endpoints) == 0 {
		return errors.New("no endpoints found")
	}

	for i, e := range c.Endpoints {
		if e.URL == "" {
			return fmt.Errorf("endpoint %d: url is required", i+1)
		}

		u, err := url.Parse(e.URL)
		if err != nil || (u.Scheme != "http" && u.Scheme != "https") || u.Host == "" {
			return fmt.Errorf("endpoint %d: invalid url %q", i+1, e.URL)
		}

		if !validMethods[strings.ToUpper(e.Method)] {
			return fmt.Errorf("endpoint %d: unsupported method %q", i+1, e.Method)
		}
	}
	return nil
}
