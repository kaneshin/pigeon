package pigeon

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectionType(t *testing.T) {
	assert := assert.New(t)

	assert.Equal("TYPE_UNSPECIFIED", DetectionType(TypeUnspecified))
	assert.Equal("FACE_DETECTION", DetectionType(FaceDetection))
	assert.Equal("LANDMARK_DETECTION", DetectionType(LandmarkDetection))
	assert.Equal("LOGO_DETECTION", DetectionType(LogoDetection))
	assert.Equal("LABEL_DETECTION", DetectionType(LabelDetection))
	assert.Equal("TEXT_DETECTION", DetectionType(TextDetection))
	assert.Equal("SAFE_SEARCH_DETECTION", DetectionType(SafeSearchDetection))
	assert.Equal("IMAGE_PROPERTIES", DetectionType(ImageProperties))
	assert.Equal("", DetectionType(-1))

	f := NewFeature(LabelDetection)
	assert.EqualValues(0, f.MaxResults)
	assert.EqualValues("LABEL_DETECTION", f.Type)
}
