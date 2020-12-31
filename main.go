package main

import (
	"github.com/Go-Amazon/actions"
	"github.com/ngs/go-amazon-product-advertising-api/amazon"
	"log"
	"os"
)

func main() {
	client, err := amazon.NewFromEnvionment()

	if err != nil {
		log.Fatal(err)
	}

	actions.HandleAction("top buttercup xxl 50", os.Stderr, client)
	// TODO: Programmatically get AWS keys and information using HTTPS instead of requiring the user to do so
	// TODO: Add a while loop for repetitive command
}
