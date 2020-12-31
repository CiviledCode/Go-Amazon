package actions

import (
	"github.com/ngs/go-amazon-product-advertising-api/amazon"
	"io"
	"log"
	"strings"
)

// Action represents a possible query that we can execute
// All actions must be registered inside of action_id.go inside of the init function
type Action interface {

	// ID represents all of the possible aliases/names that a command can be executed by
	// An ID should NOT contain any spaces. We split our input query by spaces, so any spaces in between will make the second word a param
	ID() []string

	// Usage represents the proper usage of the query
	// EX: "lookup [name] (resultsamount)"
	Usage() string

	// Execute is what executes the query. We take in writer so we can properly write necessary data to the stream we want to use
	// We also take in client to execute web queries with amazon
	// params holds all of the following strings after the ID (index 1:...)
	// We return an error so if there is ever an issue we can let the handler print it to the stream
	Execute(params []string, client *amazon.Client, writer io.Writer) error
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

	err = action.Execute(split[1:], client, writer)

	if err != nil {
		log.Fatal(err)
	}
}
