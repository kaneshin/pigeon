package pigeon

import (
	"encoding/base64"
	"io/ioutil"
	"strings"

	vision "google.golang.org/api/vision/v1"
)

var (
	emptyAnnotateImageRequest = &vision.AnnotateImageRequest{}
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

// NewBatchAnnotateImageRequest returns a pointer to a new vision's BatchAnnotateImagesRequest.
func NewBatchAnnotateImageRequest(list []string, features ...*vision.Feature) (*vision.BatchAnnotateImagesRequest, error) {
	batch := &vision.BatchAnnotateImagesRequest{}
	batch.Requests = []*vision.AnnotateImageRequest{}
	for _, v := range list {
		req, err := NewAnnotateImageRequest(v, features...)
		if err != nil {
			return nil, err
		}
		batch.Requests = append(batch.Requests, req)
	}
	return batch, nil
}

// NewAnnotateImageRequest returns a pointer to a new vision's AnnotateImagesRequest.
func NewAnnotateImageRequest(v interface{}, features ...*vision.Feature) (*vision.AnnotateImageRequest, error) {
	switch v.(type) {
	case []byte:
		// base64
		return NewAnnotateImageContentRequest(v.([]byte), features...)
	case string:
		str := v.(string)
		if strings.HasPrefix(str, "gs://") {
			// GcsImageUri: Google Cloud Storage image URI. It must be in the
			// following form:
			// "gs://bucket_name/object_name". For more
			return NewAnnotateImageSourceRequest(str, features...)
		}
		// filepath
		b, err := ioutil.ReadFile(str)
		if err != nil {
			return nil, err
		}
		return NewAnnotateImageRequest(b, features...)
	}
	return emptyAnnotateImageRequest, nil
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
