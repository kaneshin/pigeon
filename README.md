# Pigeon - Cloud Vision API on Golang

`pigeon` is a wrapper service for the Google Cloud Vision API on Golang.

## Prerequisite

You need to export a service account json file into `GOOGLE_APPLICATION_CREDENTIALS` env variable.

```
$ export GOOGLE_APPLICATION_CREDENTIALS=/path/to/service_account.json
```

Abort this execution file if you don't set the variable or unable to find the file.


## Installation

Type the below command to install if you use this application on your device.

```shell
go get github.com/kaneshin/pigeon/...
```

Make sure that `pigeon` was installed correctly:

```shell
pigeon -h
```

### Dependencies

To run this sample, you need to install this packages;

- golang.org/x/net/context
- golang.org/x/oauth2/google
- google.golang.org/api/vision/...

This sample repository is contained glide.yaml to privide `glide` command. So you should install that packages with the `glide` command.

See https://github.com/Masterminds/glide


## Usage

### CLI

```shell
go run ./cmd/pigeon lenna.jpg
go run ./cmd/pigeon gs://bucket_name/lenna.jpg
```

or if you already installed as a command.

```shell
pigeon lenna.jpg
pigeon gs://bucket_name/lenna.jpg
```

### Code

#### pigeon.Client

The `pigeon.Client` is wrapper of the `vision.Service`.

```go
// Initialize vision client by a credentials json.
client, err := pigeon.New()
if err != nil {
	panic(err)
}
```

#### vision.Feature

`vision.Feature` will be applied to `vision.AnnotateImageRequest`.

```go
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

// Choose detection types
features := []*vision.Feature{
	pigeon.NewFeature(pigeon.FaceDetection),
	pigeon.NewFeature(pigeon.LabelDetection),
	pigeon.NewFeature(pigeon.ImageProperties),
}
```

#### vision.AnnotateImageRequest

`vision.AnnotateImageRequest` needs to set the uri of the form `"gs://bucket_name/foo.png"` or byte content of image.

- Google Cloud Storage

```go
req, err := pigeon.NewAnnotateImageSourceRequest("gs://bucket_name/lenna.jpg", features...)
if err != nil {
	panic(err)
}
```

- Base64 Encoded String

```go
// e.g.) fp names "example.png"
b, err := ioutil.ReadFile(fp)
if err != nil {
	panic(err)
}
req, err = pigeon.NewAnnotateImageContentRequest(b, features...)
if err != nil {
	panic(err)
}
```

#### Submit the request to the Google Cloud Vision API

```go
// To call multiple image annotation requests.
batch := &vision.BatchAnnotateImagesRequest{
	Requests: []*vision.AnnotateImageRequest{req},
}

// Execute the "vision.images.annotate".
res, err := client.ImagesService().Annotate(batch).Do()
if err != nil {
	log.Printf("Unable to execute images annotate requests: %v\n", err)
	return 1
}
```


## Example

### input

![lenna.jpg](https://raw.githubusercontent.com/kaneshin/pigeon/master/lenna.jpg)

### output

[more detail](https://github.com/kaneshin/pigeon/blob/master/result.json)

```
[
  {
    "faceAnnotations": [
      {
        "angerLikelihood": "VERY_UNLIKELY",
        "blurredLikelihood": "VERY_UNLIKELY",
        "boundingPoly": {
          "vertices": [
            {
              "x": 143,
              "y": 43
            },
            {
              "x": 245,
              "y": 43
            },
            {
              "x": 245,
              "y": 163
            },
            {
              "x": 143,
              "y": 163
            }
          ]
        },
        "detectionConfidence": 0.99805844,
        "fdBoundingPoly": {
          "vertices": [
            {
              "x": 172,
              "y": 82
            },
            {
              "x": 241,
              "y": 82
            },
            {
              "x": 241,
              "y": 151
            },
            {
              "x": 172,
              "y": 151
            }
          ]
        },
        "headwearLikelihood": "UNLIKELY",
        "joyLikelihood": "VERY_UNLIKELY",
        "landmarkingConfidence": 0.5350582,
        "landmarks": [
          {
            "position": {
              "x": 197.90556,
              "y": 102.932,
              "z": 0.00083794753
            },
            "type": "LEFT_EYE"
          },
          {
            "position": {
              "x": 223.43489,
              "y": 102.72927,
              "z": 17.352478
            },
            "type": "RIGHT_EYE"
          },
          {
            "position": {
              "x": 189.50327,
              "y": 96.40799,
              "z": -4.1362653
            },
            "type": "LEFT_OF_LEFT_EYEBROW"
          },
          ...
```

## TODO

- More Features
- Connect to CI
- Test Code


## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)


## Author

Shintaro Kaneko <kaneshin0120@gmail.com>
