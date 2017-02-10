package credentials

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStaticCredentials(t *testing.T) {
	c := NewStaticCredentials(
		"project-id",
		"some_number",
		"-----BEGIN PRIVATE KEY-----\n....=\n-----END PRIVATE KEY-----\n",
		"visionapi@project-id.iam.gserviceaccount.com",
		"...",
	)
	creds, err := c.Get()

	assert.NoError(t, err, "Expect no error")
	assert.Equal(t, "service_account", creds.Type)
	assert.Equal(t, "project-id", creds.ProjectID)
	assert.Equal(t, "some_number", creds.PrivateKeyID)
	assert.Equal(t, "-----BEGIN PRIVATE KEY-----\n....=\n-----END PRIVATE KEY-----\n", creds.PrivateKey)
	assert.Equal(t, "visionapi@project-id.iam.gserviceaccount.com", creds.ClientEmail)
	assert.Equal(t, "...", creds.ClientID)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/auth", creds.AuthURI)
	assert.Equal(t, "https://accounts.google.com/o/oauth2/token", creds.TokenURI)
	assert.Equal(t, "https://www.googleapis.com/oauth2/v1/certs", creds.AuthProviderX509CertURL)
	assert.Equal(t, "https://www.googleapis.com/robot/v1/metadata/x509/visionapi%40project-id.iam.gserviceaccount.com", creds.ClientX509CertURL)

	t.Run("Provider", func(t *testing.T) {
		p := StaticProvider{
			Value: Value{
				ProjectID:    "project-id",
				PrivateKeyID: "some_number",
				PrivateKey:   "-----BEGIN PRIVATE KEY-----\n....=\n-----END PRIVATE KEY-----\n",
				ClientEmail:  "visionapi@project-id.iam.gserviceaccount.com",
				ClientID:     "...",
			},
		}
		creds, err := p.Retrieve()

		assert.NoError(t, err, "Expect no error")
		assert.Equal(t, "project-id", creds.ProjectID)
		assert.Equal(t, "some_number", creds.PrivateKeyID)
		assert.Equal(t, "-----BEGIN PRIVATE KEY-----\n....=\n-----END PRIVATE KEY-----\n", creds.PrivateKey)
		assert.Equal(t, "visionapi@project-id.iam.gserviceaccount.com", creds.ClientEmail)
		assert.Equal(t, "...", creds.ClientID)

		t.Run("Invalie Error", func(t *testing.T) {
			p := StaticProvider{
				Value: Value{
					ProjectID:    "",
					PrivateKeyID: "",
					PrivateKey:   "",
					ClientEmail:  "",
					ClientID:     "",
				},
			}
			creds, err := p.Retrieve()

			assert.Error(t, err, "Should be error")
			assert.Equal(t, creds, Value{})
		})
	})
}
