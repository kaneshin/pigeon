package pigeon

import (
	vision "google.golang.org/api/vision/v1"
)

func NewAnnotateImageContentRequest(body []byte, features ...*vision.Feature) (*vision.AnnotateImageRequest, error) {
	req := &vision.AnnotateImageRequest{
		Image:    NewAnnotateImageContent(body),
		Features: features,
	}
	return req, nil
}

func NewAnnotateImageSourceRequest(source string, features ...*vision.Feature) (*vision.AnnotateImageRequest, error) {
	req := &vision.AnnotateImageRequest{
		Image:    NewAnnotateImageSource(source),
		Features: features,
	}
	return req, nil
}
