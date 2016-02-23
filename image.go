package pigeon

import (
	"encoding/base64"

	vision "google.golang.org/api/vision/v1"
)

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
