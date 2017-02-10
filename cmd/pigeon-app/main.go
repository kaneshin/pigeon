package main

import (
	"flag"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/olahol/go-imageupload"
	vision "google.golang.org/api/vision/v1"

	"github.com/kaneshin/pigeon"
	"github.com/kaneshin/pigeon/cmd"
)

func main() {
	port := flag.Int("port", 0, "")
	flag.Parse()

	// Parse arguments to run this function.
	detects := cmd.DetectionsParse(flag.Args()[:])

	// Initialize vision service by a credentials json.
	client, err := pigeon.New(nil)
	if err != nil {
		panic(err)
	}

	// Initialize a router of gin.
	r := gin.Default()

	for k := range _bindata {
		html := template.Must(template.New(k).Parse(string(MustAsset(k))))
		r.SetHTMLTemplate(html)
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "views/index.html", nil)
	})

	r.POST("/upload", func(c *gin.Context) {
		img, err := imageupload.Process(c.Request, "file")
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		req, err := pigeon.NewAnnotateImageContentRequest(img.Data, detects.Features()...)
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

	if *port == 0 {
		*port = 8080
	}
	fmt.Fprintf(os.Stderr, "%v\n", r.Run(fmt.Sprintf(":%d", *port)))
}
