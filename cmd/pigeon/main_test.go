package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeatures(t *testing.T) {
	assert := assert.New(t)

	assert.False(*FaceDetection)
	assert.False(*LandmarkDetection)
	assert.False(*LogoDetection)
	assert.False(*LabelDetection)
	assert.False(*TextDetection)
	assert.False(*SafeSearchDetection)
	assert.False(*ImageProperties)

	f := features()
	if assert.Equal(1, len(f)) {
		feature := f[0]
		assert.EqualValues(0, feature.MaxResults)
		assert.Equal("FACE_DETECTION", feature.Type)
	}

	*LandmarkDetection = true
	f = features()
	if assert.Equal(1, len(f)) {
		feature := f[0]
		assert.EqualValues(0, feature.MaxResults)
		assert.Equal("LANDMARK_DETECTION", feature.Type)
	}

	*LogoDetection = true
	f = features()
	if assert.Equal(2, len(f)) {
		feature := f[0]
		assert.EqualValues(0, feature.MaxResults)
		assert.Equal("LANDMARK_DETECTION", feature.Type)
		feature = f[1]
		assert.EqualValues(0, feature.MaxResults)
		assert.Equal("LOGO_DETECTION", feature.Type)
	}
}
