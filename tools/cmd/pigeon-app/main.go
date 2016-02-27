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
	"github.com/kaneshin/pigeon/tools/cmd"
)

func main() {
	port := flag.Int("port", 0, "")
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		// Parse arguments to run this function.
		detects := cmd.DetectionsParse([]string{})
		detects.Usage()
		os.Exit(1)
	}

	// Parse arguments to run this function.
	detects := cmd.DetectionsParse(flag.Args()[1:])

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
