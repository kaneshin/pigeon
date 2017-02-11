package main

import (
	"flag"

	vision "google.golang.org/api/vision/v1"

	"github.com/kaneshin/pigeon"
)

const (
	defaultDetection = pigeon.LabelDetection
)

// Detections type
type Detections struct {
	face            bool
	landmark        bool
	logo            bool
	label           bool
	text            bool
	safeSearch      bool
	imageProperties bool
	args            []string
	flag            *flag.FlagSet
}

// DetectionsParse parses the command-line flags from arguments and returns
// a new pointer of a Detections object..
func DetectionsParse(args []string) *Detections {
	f := flag.NewFlagSet("Detections", flag.ExitOnError)
	faceDetection := f.Bool("face", false, "This flag specifies the face detection of the feature")
	landmarkDetection := f.Bool("landmark", false, "This flag specifies the landmark detection of the feature")
	logoDetection := f.Bool("logo", false, "This flag specifies the logo detection of the feature")
	labelDetection := f.Bool("label", false, "This flag specifies the label detection of the feature")
	textDetection := f.Bool("text", false, "This flag specifies the text detection (OCR) of the feature")
	safeSearchDetection := f.Bool("safe-search", false, "This flag specifies the safe-search of the feature")
	imageProperties := f.Bool("image-properties", false, "This flag specifies the image safe-search properties of the feature")
	f.Usage = func() {
		f.PrintDefaults()
	}
	f.Parse(args)
	return &Detections{
		face:            *faceDetection,
		landmark:        *landmarkDetection,
		logo:            *logoDetection,
		label:           *labelDetection,
		text:            *textDetection,
		safeSearch:      *safeSearchDetection,
		imageProperties: *imageProperties,
		flag:            f,
	}
}

// Args returns the non-flag command-line arguments.
func (d *Detections) Args() []string {
	return d.flag.Args()
}

// Usage prints options of the Detection object.
func (d *Detections) Usage() {
	d.flag.Usage()
}

// Features returns a slice of pointers of vision.Feature.
func (d *Detections) Features() []*vision.Feature {
	list := []int{}
	if d.face {
		list = append(list, pigeon.FaceDetection)
	}
	if d.landmark {
		list = append(list, pigeon.LandmarkDetection)
	}
	if d.logo {
		list = append(list, pigeon.LogoDetection)
	}
	if d.label {
		list = append(list, pigeon.LabelDetection)
	}
	if d.text {
		list = append(list, pigeon.TextDetection)
	}
	if d.safeSearch {
		list = append(list, pigeon.SafeSearchDetection)
	}
	if d.imageProperties {
		list = append(list, pigeon.ImageProperties)
	}

	if len(list) == 0 {
		list = append(list, defaultDetection)
	}

	features := make([]*vision.Feature, len(list))
	for i := 0; i < len(list); i++ {
		features[i] = pigeon.NewFeature(list[i])
	}
	return features
}
