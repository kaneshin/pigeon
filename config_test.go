package pigeon

import (
	"net/http"
	"testing"

	"github.com/kaneshin/pigeon/credentials"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	conf := NewConfig()
	assert.NotNil(conf)
	assert.Nil(conf.Credentials)
	assert.Nil(conf.HTTPClient)

	creds := credentials.NewApplicationCredentials("")
	client := http.DefaultClient
	conf.WithCredentials(creds).
		WithHTTPClient(client)
	assert.NotNil(conf.Credentials)
	assert.NotNil(conf.HTTPClient)
}
