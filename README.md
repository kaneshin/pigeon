# Pigeon - Google Cloud Vision API on Golang

`pigeon` is a service for the Google Cloud Vision API on Golang.

## Badges

[![wercker status](https://app.wercker.com/status/265bd30a85f806655926be3ded5eff13/s "wercker status")](https://app.wercker.com/project/bykey/265bd30a85f806655926be3ded5eff13)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](http://makeapullrequest.com)
[![Code Climate](https://codeclimate.com/github/kaneshin/pigeon/badges/gpa.svg)](https://codeclimate.com/github/kaneshin/pigeon)


## Prerequisite

You need to export a service account json file to `GOOGLE_APPLICATION_CREDENTIALS` variable.

```
$ export GOOGLE_APPLICATION_CREDENTIALS=/path/to/service_account.json
```


## Installation

Type the following command to install `pigeon`.

```shell
$ go get github.com/kaneshin/pigeon/...
```

Make sure that `pigeon` was installed correctly:

```shell
$ pigeon -h
```


## Usage

### Command

```shell
$ pigeon assets/lenna.jpg
$ pigeon gs://bucket_name/lenna.jpg
```

`pigeon` is available to submit request with external image source (i.e. Google Cloud Storage image location).


### `pigeon` package

```go
import "github.com/kaneshin/pigeon"
```

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
src := "gs://bucket_name/lenna.jpg"
req, err := pigeon.NewAnnotateImageSourceRequest(src, features...)
if err != nil {
	panic(err)
}
```

- Base64 Encoded String

```go
b, err := ioutil.ReadFile(filename)
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
batch, err := pigeon.NewBatchAnnotateImageRequest(list, features()...)
if err != nil {
	panic(err)
}

// Execute the "vision.images.annotate".
res, err := client.ImagesService().Annotate(batch).Do()
if err != nil {
	panic(err)
}
```


## Example

### Pigeon

![pigeon](https://raw.githubusercontent.com/kaneshin/pigeon/master/assets/pigeon.png)

#### input

```shell
$ pigeon -label assets/pigeon.png
```

#### output

```json
[
  {
    "labelAnnotations": [
      {
        "description": "bird",
        "mid": "/m/015p6",
        "score": 0.825656
      },
      {
        "description": "anatidae",
        "mid": "/m/01c_0l",
        "score": 0.58264238
      }
    ]
  }
]
```


### Lenna

![lenna](https://raw.githubusercontent.com/kaneshin/pigeon/master/assets/lenna.jpg)

#### input

```shell
$ pigeon -safe-search assets/lenna.jpg
```

#### output

```json
[
  {
    "safeSearchAnnotation": {
      "adult": "POSSIBLE",
      "medical": "UNLIKELY",
      "spoof": "VERY_UNLIKELY",
      "violence": "VERY_UNLIKELY"
    }
  }
]
```


## License

[The MIT License (MIT)](http://kaneshin.mit-license.org/)


## Author

Shintaro Kaneko <kaneshin0120@gmail.com>
