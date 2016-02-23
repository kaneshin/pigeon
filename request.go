package pigeon

import (
	vision "google.golang.org/api/vision/v1"
	"io/ioutil"
	"strings"
)

var (
	emptyAnnotateImageRequest = &vision.AnnotateImageRequest{}
)

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

func NewAnnotateImageRequest(v interface{}, features ...*vision.Feature) (*vision.AnnotateImageRequest, error) {
	switch v.(type) {
	case []byte:
		// base64
		return NewAnnotateImageContentRequest(v.([]byte), features...)
	case string:
		str := v.(string)
		if strings.HasPrefix(str, "gs://") {
			// gs://bucket-name
			return NewAnnotateImageSourceRequest(str, features...)
		} else {
			// filepath
			b, err := ioutil.ReadFile(str)
			if err != nil {
				return nil, err
			}
			return NewAnnotateImageRequest(b, features...)
		}
	}
	return emptyAnnotateImageRequest, nil
}

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
