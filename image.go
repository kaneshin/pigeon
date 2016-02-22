package pigeon

import (
	"encoding/base64"

	vision "google.golang.org/api/vision/v1"
)

// NewAnnotateImageContent: Image content, represented as a stream of bytes.
func NewAnnotateImageContent(body []byte) *vision.Image {
	return &vision.Image{
		Content: base64.StdEncoding.EncodeToString(body),
	}
}

// NewAnnotateImageSource: External image source (i.e. Google Cloud Storage image
// location).
func NewAnnotateImageSource(source string) *vision.Image {
	return &vision.Image{
		Source: &vision.ImageSource{
			GcsImageUri: source,
		},
	}
}
