package pigeon

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/oauth2/google"
	vision "google.golang.org/api/vision/v1"

	"github.com/kaneshin/pigeon/credentials"
)

type (
	// A Client provides cloud vision service.
	Client struct {
		// The context object to use when signing requests.
		// Defaults to `context.Background()`.
		// context context.Context

		// The Config provides service configuration for service clients.
		config *Config

		// The service object.
		service *vision.Service
	}
)

// New returns a pointer to a new Client object.
func New(c *Config, httpClient ...*http.Client) (*Client, error) {
	if c == nil {
		// Sets a configuration if passed nil value.
		c = NewConfig()
	}

	// Use HTTP Client if assigned
	if len(httpClient) > 0 {
		srv, err := vision.New(httpClient[0])
		if err != nil {
			return nil, fmt.Errorf("Unable to retrieve vision Client %v", err)
		}
		return &Client{
			config:  c,
			service: srv,
		}, nil
	}

	if c.Credentials == nil {
		// Sets application credentials by defaults.
		c.Credentials = credentials.NewApplicationCredentials("")
	}

	creds, err := c.Credentials.Get()
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(creds)

	config, err := google.JWTConfigFromJSON(b, vision.CloudPlatformScope)
	if err != nil {
		return nil, fmt.Errorf("Unable to parse client secret file to config: %v", err)
	}
	client := config.Client(context.Background())

	srv, err := vision.New(client)
	if err != nil {
		return nil, fmt.Errorf("Unable to retrieve vision Client %v", err)
	}

	return &Client{
		config:  c,
		service: srv,
	}, nil
}

// ImagesService returns a pointer to a vision's ImagesService object.
func (c Client) ImagesService() *vision.ImagesService {
	return c.service.Images
}

// NewBatchAnnotateImageRequest returns a pointer to a new vision's BatchAnnotateImagesRequest.
func (c Client) NewBatchAnnotateImageRequest(list []string, features ...*vision.Feature) (*vision.BatchAnnotateImagesRequest, error) {
	batch := &vision.BatchAnnotateImagesRequest{}
	batch.Requests = []*vision.AnnotateImageRequest{}
	for _, v := range list {
		req, err := c.NewAnnotateImageRequest(v, features...)
		if err != nil {
			return nil, err
		}
		batch.Requests = append(batch.Requests, req)
	}
	return batch, nil
}

// NewAnnotateImageRequest returns a pointer to a new vision's AnnotateImagesRequest.
func (c Client) NewAnnotateImageRequest(v interface{}, features ...*vision.Feature) (*vision.AnnotateImageRequest, error) {
	switch v.(type) {
	case []byte:
		// base64
		return NewAnnotateImageContentRequest(v.([]byte), features...)
	case string:
		u, err := url.Parse(v.(string))
		if err != nil {
			return nil, err
		}
		switch u.Scheme {
		case "gs":
			// GcsImageUri: Google Cloud Storage image URI. It must be in the
			// following form:
			// "gs://bucket_name/object_name". For more
			return NewAnnotateImageSourceRequest(u.String(), features...)
		case "http", "https":
			httpClient := c.config.HTTPClient
			if httpClient == nil {
				httpClient = http.DefaultClient
			}
			resp, err := httpClient.Get(u.String())
			if err != nil {
				return nil, err
			}
			defer resp.Body.Close()
			if resp.StatusCode >= http.StatusBadRequest {
				return nil, http.ErrMissingFile
			}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			return c.NewAnnotateImageRequest(body, features...)
		}
		// filepath
		b, err := ioutil.ReadFile(v.(string))
		if err != nil {
			return nil, err
		}
		return c.NewAnnotateImageRequest(b, features...)
	}
	return &vision.AnnotateImageRequest{}, nil
}

// NewAnnotateImageContentRequest returns a pointer to a new vision's AnnotateImagesRequest.
func NewAnnotateImageContentRequest(body []byte, features ...*vision.Feature) (*vision.AnnotateImageRequest, error) {
	req := &vision.AnnotateImageRequest{
		Image:    NewAnnotateImageContent(body),
		Features: features,
	}
	return req, nil
}

// NewAnnotateImageSourceRequest returns a pointer to a new vision's AnnotateImagesRequest.
func NewAnnotateImageSourceRequest(source string, features ...*vision.Feature) (*vision.AnnotateImageRequest, error) {
	req := &vision.AnnotateImageRequest{
		Image:    NewAnnotateImageSource(source),
		Features: features,
	}
	return req, nil
}

// NewAnnotateImageContent returns a pointer to a new vision's Image.
// It's contained image content, represented as a stream of bytes.
func NewAnnotateImageContent(body []byte) *vision.Image {
	return &vision.Image{
		// Content: Image content, represented as a stream of bytes.
		Content: base64.StdEncoding.EncodeToString(body),
	}
}

// NewAnnotateImageSource returns a pointer to a new vision's Image.
// It's contained external image source (i.e. Google Cloud Storage image
// location).
func NewAnnotateImageSource(source string) *vision.Image {
	return &vision.Image{
		Source: &vision.ImageSource{
			GcsImageUri: source,
		},
	}
}
