package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	vision "google.golang.org/api/vision/v1"
	"net/http"

	"github.com/kaneshin/pigeon"
)

var (
	faceDetection       = false
	landmarkDetection   = false
	logoDetection       = false
	labelDetection      = false
	textDetection       = false
	safeSearchDetection = false
	imageProperties     = false
)

var defaultDetection = pigeon.LabelDetection

func features() []*vision.Feature {
	list := []int{}
	if faceDetection {
		list = append(list, pigeon.FaceDetection)
	}
	if landmarkDetection {
		list = append(list, pigeon.LandmarkDetection)
	}
	if logoDetection {
		list = append(list, pigeon.LogoDetection)
	}
	if labelDetection {
		list = append(list, pigeon.LabelDetection)
	}
	if textDetection {
		list = append(list, pigeon.TextDetection)
	}
	if safeSearchDetection {
		list = append(list, pigeon.SafeSearchDetection)
	}
	if imageProperties {
		list = append(list, pigeon.ImageProperties)
	}

	// Default
	if len(list) == 0 {
		list = append(list, defaultDetection)
	}

	features := make([]*vision.Feature, len(list))
	for i := 0; i < len(list); i++ {
		features[i] = pigeon.NewFeature(list[i])
	}
	return features
}

func main() {
	port := flag.Int("port", 0, "")
	flag.Parse()
	if *port == 0 {
		*port = 8080
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.File("index.html")
	})

	r.POST("/upload", func(c *gin.Context) {
		client, err := pigeon.New()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		img, err := imageupload.Process(c.Request, "file")
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		req, err := pigeon.NewAnnotateImageContentRequest(img.Data, features()...)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		batch := &vision.BatchAnnotateImagesRequest{
			Requests: []*vision.AnnotateImageRequest{req},
		}

		res, err := client.ImagesService().Annotate(batch).Do()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, res)
	})

	r.Run(fmt.Sprintf(":%d", *port))
}
