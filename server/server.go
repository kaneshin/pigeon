package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	vision "google.golang.org/api/vision/v1"

	"github.com/kaneshin/pigeon"
	"github.com/kaneshin/pigeon/harness"
)

func main() {
	port := flag.Int("port", 0, "")
	harness.FlagParse()
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

		req, err := pigeon.NewAnnotateImageContentRequest(img.Data, harness.Features()...)
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
