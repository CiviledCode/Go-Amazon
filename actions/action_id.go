package actions

import (
	"fmt"
	"strings"
)

var RegisteredActions = map[string]func() Action{}

// FromID() returns the proper action based on it's ID
func FromID(ID string) (error, Action) {
	for id, action := range RegisteredActions {
		if strings.ToLower(ID) == id {
			return nil, action()
		}
	}

	return fmt.Errorf("couldn't find action under the id %s", ID), nil
}

func init() {
	// Push all of the actions into a list
	actions := []func() Action{
		func() Action { return &ProductSearch{} },
	}

	// Register all actions from our list into a map containing aliases too
	for _, action := range actions {
		for _, aliase := range action().ID() {
			RegisteredActions[aliase] = action
		}
	}
}
