package pigeon

import (
	vision "google.golang.org/api/vision/v1"
)

const (
	TypeUnspecified = iota + 1
	FaceDetection
	LandmarkDetection
	LogoDetection
	LabelDetection
	TextDetection
	SafeSearchDetection
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

func NewFeature(d int) *vision.Feature {
	return &vision.Feature{Type: DetectionType(d)}
}
