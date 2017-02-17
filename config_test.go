package pigeon

import (
	"net/http"
	"testing"

	"github.com/kaneshin/pigeon/credentials"
	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	cfg := NewConfig()
	assert.NotNil(cfg)
	assert.Nil(cfg.Credentials)
	assert.Nil(cfg.HTTPClient)

	creds := credentials.NewApplicationCredentials("")
	client := http.DefaultClient
	cfg.WithCredentials(creds).
		WithHTTPClient(client)
	assert.NotNil(cfg.Credentials)
	assert.NotNil(cfg.HTTPClient)
}

func Benchmark_Config(b *testing.B) {
	var c Config

	b.Run("NewConfig", func(b *testing.B) {
		b.ReportAllocs()
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			c = NewConfig()
		}
	})
}
