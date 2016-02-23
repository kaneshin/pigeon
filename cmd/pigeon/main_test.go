package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFeatures(t *testing.T) {
	assert := assert.New(t)

	assert.False(*faceDetection)
	assert.False(*landmarkDetection)
	assert.False(*logoDetection)
	assert.False(*labelDetection)
	assert.False(*textDetection)
	assert.False(*safeSearchDetection)
	assert.False(*imageProperties)

	f := features()
	if assert.Equal(1, len(f)) {
		feature := f[0]
		assert.EqualValues(0, feature.MaxResults)
		assert.Equal("LABEL_DETECTION", feature.Type)
	}

	*landmarkDetection = true
	f = features()
	if assert.Equal(1, len(f)) {
		feature := f[0]
		assert.EqualValues(0, feature.MaxResults)
		assert.Equal("LANDMARK_DETECTION", feature.Type)
	}

	*logoDetection = true
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
