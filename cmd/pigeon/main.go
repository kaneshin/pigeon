package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kaneshin/pigeon"
	"github.com/kaneshin/pigeon/harness"
)

func main() {
	// Parse arguments to run this function.
	harness.FlagParse()

	if args := harness.FlagArgs(); len(args) == 0 {
		fmt.Fprintf(os.Stderr, "pigeon [options] <source>\n")
		os.Exit(1)
	}

	// Initialize vision service by a credentials json.
	client, err := pigeon.New()
	if err != nil {
		log.Fatalf("Unable to retrieve vision service: %v\n", err)
	}

	// To call multiple image annotation requests.
	batch, err := pigeon.NewBatchAnnotateImageRequest(harness.FlagArgs(), harness.Features()...)
	if err != nil {
		log.Fatalf("Unable to retrieve image request: %v\n", err)
	}

	// Execute the "vision.images.annotate".
	res, err := client.ImagesService().Annotate(batch).Do()
	if err != nil {
		log.Fatalf("Unable to execute images annotate requests: %v\n", err)
	}

	// Marshal annotations from responses
	body, err := json.MarshalIndent(res.Responses, "", "  ")
	if err != nil {
		log.Fatalf("Unable to marshal the response: %v\n", err)
	}
	fmt.Println(string(body))
}
