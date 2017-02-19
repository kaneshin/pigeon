package pigeon

import (
	"net/http"

	"github.com/kaneshin/pigeon/credentials"
)

type (
	// A Config provides service configuration for service clients. By default,
	// all clients will use the {defaults.DefaultConfig} structure.
	Config struct {
		// The credentials object to use when signing requests.
		// Defaults to application credentials file.
		Credentials *credentials.Credentials

		// The HTTP client to use when sending requests.
		// Defaults to `http.DefaultClient`.
		HTTPClient *http.Client
	}
)

// NewConfig returns a new pointer Config object.
func NewConfig() *Config {
	return &Config{}
}

// WithCredentials sets a config Credentials value returning a Config pointer
// for chaining.
func (c *Config) WithCredentials(creds *credentials.Credentials) *Config {
	c.Credentials = creds
	return c
}

// WithHTTPClient sets a config HTTPClient value returning a Config pointer
// for chaining.
func (c *Config) WithHTTPClient(client *http.Client) *Config {
	c.HTTPClient = client
	return c
}
