package pigeon

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	vision "google.golang.org/api/vision/v1"
)

func TestClient(t *testing.T) {
	os.Clearenv()

	assert := assert.New(t)

	cfg := NewConfig()
	client, err := New(cfg)
	assert.Nil(client)
	assert.Error(err)

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "credentials/example.json")
	cfg = NewConfig()
	client, err = New(cfg)
	assert.NotNil(client)
	assert.NoError(err)
	assert.NotNil(client.service)
	assert.NotNil(client.ImagesService())
}

func TestNewAnnotateImageRequest(t *testing.T) {
	assert := assert.New(t)

	var (
		req *vision.AnnotateImageRequest
		err error
	)
	const (
		gcsImageURI      = "gs://bucket/sample.png"
		fp               = "assets/lenna.jpg"
		imageURI         = "https://httpbin.org/image/jpeg"
		imageURINoExists = "https://httpbin.org/image/jpeg/none"
	)
	features := NewFeature(LabelDetection)
	client, err := New(nil)
	assert.NoError(err)

	// GCS
	req, err = client.NewAnnotateImageRequest(gcsImageURI, features)
	assert.NoError(err)
	if assert.NotNil(req) {
		assert.Equal("", req.Image.Content)
		assert.Equal(gcsImageURI, req.Image.Source.GcsImageUri)
	}

	// Filepath
	req, err = client.NewAnnotateImageRequest(fp, features)
	assert.NoError(err)
	if assert.NotNil(req) {
		assert.NotEqual("", req.Image.Content)
		assert.Nil(req.Image.Source)
	}

	// Image URI
	req, err = client.NewAnnotateImageRequest(imageURI, features)
	assert.NoError(err)
	if assert.NotNil(req) {
		assert.NotEqual("", req.Image.Content)
		assert.Nil(req.Image.Source)
	}

	req, err = client.NewAnnotateImageRequest(imageURINoExists, features)
	assert.Error(err)
	assert.Nil(req)
}
