package credentials

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplicationCredentials(t *testing.T) {
	os.Clearenv()

	c := NewApplicationCredentials("example.json")
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
		os.Clearenv()

		p := ApplicationCredentialsProvider{Filename: "example.json"}
		creds, err := p.Retrieve()

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

		t.Run("No File Error", func(t *testing.T) {
			p := ApplicationCredentialsProvider{Filename: "notexist.json"}
			creds, err := p.Retrieve()

			assert.Error(t, err, "Should be error")
			assert.Equal(t, creds, Value{})
		})

		t.Run("Not JSON File Error", func(t *testing.T) {
			p := ApplicationCredentialsProvider{Filename: "credentials.go"}
			creds, err := p.Retrieve()

			assert.Error(t, err, "Should be error")
			assert.Equal(t, creds, Value{})
		})
	})

	t.Run("ProviderWithGOOGLE_APPLICATION_CREDENTIALS_FILE", func(t *testing.T) {
		os.Clearenv()

		p := ApplicationCredentialsProvider{}
		creds, err := p.Retrieve()
		assert.Error(t, err, "Should be error")

		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "example.json")
		creds, err = p.Retrieve()

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
	})
}

func BenchmarkApplicationCredentialsProvider(b *testing.B) {
	os.Clearenv()

	p := ApplicationCredentialsProvider{Filename: "example.json"}
	_, err := p.Retrieve()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := p.Retrieve()
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
