package pigeon

import (
	vision "google.golang.org/api/vision/v1"
)

const (
	// TypeUnspecified - Unspecified feature type.
	TypeUnspecified = iota
	// FaceDetection - Run face detection.
	FaceDetection
	// LandmarkDetection - Run landmark detection.
	LandmarkDetection
	// LogoDetection - Run logo detection.
	LogoDetection
	// LabelDetection - Run label detection.
	LabelDetection
	// TextDetection - Run OCR.
	TextDetection
	// SafeSearchDetection - Run various computer vision models to
	SafeSearchDetection
	// ImageProperties - compute image safe-search properties.
	ImageProperties
)

// DetectionType returns a value of detection type.
func DetectionType(d int) string {
	switch d {
	case TypeUnspecified:
		return "TYPE_UNSPECIFIED"
	case FaceDetection:
		return "FACE_DETECTION"
	case LandmarkDetection:
		return "LANDMARK_DETECTION"
	case LogoDetection:
		return "LOGO_DETECTION"
	case LabelDetection:
		return "LABEL_DETECTION"
	case TextDetection:
		return "TEXT_DETECTION"
	case SafeSearchDetection:
		return "SAFE_SEARCH_DETECTION"
	case ImageProperties:
		return "IMAGE_PROPERTIES"
	}
	return ""
}

// NewFeature returns a pointer to a new vision's Feature object.
func NewFeature(d int) *vision.Feature {
	return &vision.Feature{Type: DetectionType(d)}
}
