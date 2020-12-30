package actions

import (
	"github.com/ngs/go-amazon-product-advertising-api/amazon"
	"io"
	"log"
	"strings"
)

type Action interface {
	ID() []string

	Execute(params []string, client *amazon.Client) error
}

/*
	HandleAction() takes in a string, determines the action type, and executes it

	We take in writer so we can have file, console, and stream write options for scalability
*/
func HandleAction(s string, writer io.Writer, client *amazon.Client) {
	log.SetOutput(writer)

	split := strings.Split(s, " ")

	err, action := FromID(split[0])

	if err != nil {
		log.Fatal(err)
	}

	err = action.Execute(split[1:], client)

	if err != nil {
		log.Fatal(err)
	}
}
