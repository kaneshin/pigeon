package credentials

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type stubProvider struct {
	creds Value
	err   error
}

func (s *stubProvider) Retrieve() (Value, error) {
	return s.creds, s.err
}

func TestCredentialsGet(t *testing.T) {
	c := NewCredentials(&stubProvider{
		creds: Value{
			Type:                    "service_account",
			ProjectID:               "project-id",
			PrivateKeyID:            "some_number",
			PrivateKey:              "-----BEGIN PRIVATE KEY-----\n....=\n-----END PRIVATE KEY-----\n",
			ClientEmail:             "visionapi@project-id.iam.gserviceaccount.com",
			ClientID:                "...",
			AuthURI:                 "https://accounts.google.com/o/oauth2/auth",
			TokenURI:                "https://accounts.google.com/o/oauth2/token",
			AuthProviderX509CertURL: "https://www.googleapis.com/oauth2/v1/certs",
			ClientX509CertURL:       "https://www.googleapis.com/robot/v1/metadata/x509/visionapi%40project-id.iam.gserviceaccount.com",
		},
	})

	creds, err := c.Get()

	assert := assert.New(t)
	assert.NoError(err, "Expected no error")
	assert.True(creds.IsValid())
	assert.Equal("service_account", creds.Type)
	assert.Equal("project-id", creds.ProjectID)
	assert.Equal("some_number", creds.PrivateKeyID)
	assert.Equal("-----BEGIN PRIVATE KEY-----\n....=\n-----END PRIVATE KEY-----\n", creds.PrivateKey)
	assert.Equal("visionapi@project-id.iam.gserviceaccount.com", creds.ClientEmail)
	assert.Equal("...", creds.ClientID)
	assert.Equal("https://accounts.google.com/o/oauth2/auth", creds.AuthURI)
	assert.Equal("https://accounts.google.com/o/oauth2/token", creds.TokenURI)
	assert.Equal("https://www.googleapis.com/oauth2/v1/certs", creds.AuthProviderX509CertURL)
	assert.Equal("https://www.googleapis.com/robot/v1/metadata/x509/visionapi%40project-id.iam.gserviceaccount.com", creds.ClientX509CertURL)
}

func TestCredentialsGetWithError(t *testing.T) {
	c := NewCredentials(&stubProvider{err: errors.New("provider error")})

	v, err := c.Get()
	assert := assert.New(t)
	assert.Error(err)
	assert.False(v.IsValid())
}
