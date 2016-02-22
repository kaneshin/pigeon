package pigeon

import (
	"encoding/json"
	"fmt"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	vision "google.golang.org/api/vision/v1"

	"github.com/kaneshin/pigeon/credentials"
)

type (
	Client struct {
		service *vision.Service
	}
)

// New returns a pointer to a new Client object.
func New() (*Client, error) {
	c := credentials.NewApplicationCredentials("")
	creds, err := c.Get()
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
		service: srv,
	}, nil
}

func (c Client) ImagesService() *vision.ImagesService {
	return c.service.Images
}
