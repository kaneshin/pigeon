package app

import (
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	vision "google.golang.org/api/vision/v1"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"

	"github.com/gin-gonic/gin"
	"github.com/kaneshin/pigeon"
	"github.com/olahol/go-imageupload"
)

func init() {
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
		ctx := appengine.NewContext(c.Request)

		img, err := imageupload.Process(c.Request, "file")
		if err != nil {
			msg := fmt.Sprintf("Unable to upload the image: %v", err)
			log.Errorf(ctx, msg)
			c.JSON(http.StatusBadRequest, msg)
			return
		}

		// Initialize vision service
		httpClient := &http.Client{
			Transport: &oauth2.Transport{
				Source: google.AppEngineTokenSource(ctx, vision.CloudPlatformScope),
				Base:   &urlfetch.Transport{Context: ctx},
			},
		}

		config := pigeon.NewConfig().WithHTTPClient(urlfetch.Client(ctx))
		client, err := pigeon.New(config, httpClient)
		if err != nil {
			msg := fmt.Sprintf("Unable to retrieve vision service: %v", err)
			log.Errorf(ctx, msg)
			c.JSON(http.StatusBadRequest, msg)
			return
		}

		feature := pigeon.NewFeature(pigeon.LabelDetection)
		req, err := pigeon.NewAnnotateImageContentRequest(img.Data, feature)
		if err != nil {
			msg := fmt.Sprintf("Unable to retrieve image request: %v", err)
			log.Errorf(ctx, msg)
			c.JSON(http.StatusBadRequest, msg)
			return
		}

		batch := &vision.BatchAnnotateImagesRequest{
			Requests: []*vision.AnnotateImageRequest{req},
		}

		// Execute the "vision.images.annotate".
		res, err := client.ImagesService().Annotate(batch).Do()
		if err != nil {
			msg := fmt.Sprintf("Unable to execute images annotate requests: %v", err)
			log.Errorf(ctx, msg)
			c.JSON(http.StatusInternalServerError, msg)
			return
		}

		c.JSON(http.StatusOK, res)
	})

	http.Handle("/", r)
}
