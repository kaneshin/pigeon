package pigeon

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	os.Clearenv()

	assert := assert.New(t)

	client, err := New()
	assert.Nil(client)
	assert.Error(err)

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "credentials/example.json")
	client, err = New()
	assert.NotNil(client)
	assert.NoError(err)
	assert.NotNil(client.service)
	assert.NotNil(client.ImagesService())
}
