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

	assert := assert.New(t)
	assert.NoError(err, "Expect no error")

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

func TestApplicationCredentialsProvider(t *testing.T) {
	os.Clearenv()

	p := ApplicationCredentialsProvider{Filename: "example.json"}
	creds, err := p.Retrieve()

	assert := assert.New(t)
	assert.NoError(err, "Expect no error")

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

func TestApplicationCredentialsProviderWithGOOGLE_APPLICATION_CREDENTIALS_FILE(t *testing.T) {
	os.Clearenv()

	assert := assert.New(t)
	p := ApplicationCredentialsProvider{}
	creds, err := p.Retrieve()
	assert.Error(err, "Should be error")

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "example.json")
	creds, err = p.Retrieve()

	assert.NoError(err, "Expect no error")

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
