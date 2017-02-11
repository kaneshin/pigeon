package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetections(t *testing.T) {
	assert := assert.New(t)

	var (
		args    = []string{}
		detects *Detections
	)

	detects = DetectionsParse(args)
	assert.EqualValues(0, len(detects.Args()))
	f := detects.Features()
	if assert.EqualValues(1, len(f)) {
		feature := f[0]
		assert.EqualValues(0, feature.MaxResults)
		assert.Equal("LABEL_DETECTION", feature.Type)
	}
	assert.False(detects.face)
	assert.False(detects.landmark)
	assert.False(detects.logo)
	assert.False(detects.label)
	assert.False(detects.text)
	assert.False(detects.safeSearch)
	assert.False(detects.imageProperties)

	args = []string{"-landmark", "lenna.jpg"}
	detects = DetectionsParse(args)
	assert.EqualValues(1, len(detects.Args()))
	f = detects.Features()
	if assert.Equal(1, len(f)) {
		feature := f[0]
		assert.EqualValues(0, feature.MaxResults)
		assert.Equal("LANDMARK_DETECTION", feature.Type)
	}
	assert.False(detects.face)
	assert.True(detects.landmark)
	assert.False(detects.logo)
	assert.False(detects.label)
	assert.False(detects.text)
	assert.False(detects.safeSearch)
	assert.False(detects.imageProperties)

	args = []string{"-face", "-landmark", "-logo", "-label", "-text", "-safe-search", "-image-properties", "lenna.jpg"}
	detects = DetectionsParse(args)
	assert.EqualValues(1, len(detects.Args()))
	assert.EqualValues(7, len(detects.Features()))
	assert.True(detects.face)
	assert.True(detects.landmark)
	assert.True(detects.logo)
	assert.True(detects.label)
	assert.True(detects.text)
	assert.True(detects.safeSearch)
	assert.True(detects.imageProperties)
}
