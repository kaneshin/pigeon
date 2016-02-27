package pigeon

import (
	"os"
	"testing"

	"github.com/kaneshin/pigeon/credentials"
	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	os.Clearenv()

	assert := assert.New(t)

	creds := credentials.NewApplicationCredentials("")
	client, err := New(creds)
	assert.Nil(client)
	assert.Error(err)

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "credentials/example.json")
	creds = credentials.NewApplicationCredentials("")
	client, err = New(creds)
	assert.NotNil(client)
	assert.NoError(err)
	assert.NotNil(client.service)
	assert.NotNil(client.ImagesService())
}
