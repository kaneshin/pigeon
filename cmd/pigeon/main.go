package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	vision "google.golang.org/api/vision/v1"

	"github.com/kaneshin/pigeon"
)

func run() int {
	// Parse arguments to run this function.
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		log.Printf("Argument is required.")
		return 1
	}

	// Initialize vision service by a credentials json.
	client, err := pigeon.New()
	if err != nil {
		log.Printf("Unable to retrieve vision service: %v\n", err)
		return 1
	}

	features := []*vision.Feature{
		pigeon.NewFeature(pigeon.FaceDetection),
	}

	// Create request by given argument (last one).
	fp := args[len(args)-1]
	var req *vision.AnnotateImageRequest
	if strings.HasPrefix(fp, "gs://") {
		// gs://bucket-name
		req, err = pigeon.NewAnnotateImageSourceRequest(fp, features...)
		if err != nil {
			log.Printf("Unable to retrieve image request: %v\n", err)
			return 1
		}
	} else {
		// base64
		b, err := ioutil.ReadFile(fp)
		if err != nil {
			log.Printf("Unable to read an image by file path: %v\n", err)
			return 1
		}
		req, err = pigeon.NewAnnotateImageContentRequest(b, features...)
		if err != nil {
			log.Printf("Unable to retrieve image request: %v\n", err)
			return 1
		}
	}

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

	// Marshal annotations from responses
	body, err := json.MarshalIndent(res.Responses, "", "  ")
	if err != nil {
		log.Printf("Unable to marshal the response: %v\n", err)
		return 1
	}
	fmt.Println(string(body))

	return 0
}

func main() {
	os.Exit(run())
}
