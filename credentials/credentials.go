// Package credentials provides credential retrieval and management
package credentials

import (
	"sync"
)

// A Value is the service account credentials value for individual credential fields.
type Value struct {
	Type                    string `json:"type"`
	ProjectID               string `json:"project_id"`
	PrivateKeyID            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientID                string `json:"client_id"`
	AuthURI                 string `json:"auth_uri"`
	TokenURI                string `json:"token_uri"`
	AuthProviderX509CertURL string `json:"auth_provider_x509_cert_url"`
	ClientX509CertURL       string `json:"client_x509_cert_url"`
}

// IsValid ...
func (v *Value) IsValid() bool {
	return v.ProjectID != "" &&
		v.PrivateKeyID != "" &&
		v.PrivateKey != "" &&
		v.ClientEmail != "" &&
		v.ClientID != ""
}

// A Provider is the interface for any component which will provide credentials
// Value.
type Provider interface {
	// Refresh returns nil if it successfully retrieved the value.
	// Error is returned if the value were not obtainable, or empty.
	Retrieve() (Value, error)
}

// A Credentials provides synchronous safe retrieval of service account
// credentials Value.
type Credentials struct {
	creds    Value
	m        sync.Mutex
	provider Provider
}

// NewCredentials returns a pointer to a new Credentials with the provider set.
func NewCredentials(provider Provider) *Credentials {
	return &Credentials{
		provider: provider,
	}
}

// Get returns the credentials value, or error if the credentials Value failed
// to be retrieved.
func (c *Credentials) Get() (Value, error) {
	c.m.Lock()
	defer c.m.Unlock()
	return c.provider.Retrieve()
}
