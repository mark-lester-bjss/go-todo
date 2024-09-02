package cliCommand

import (
	"fmt"
	"toDoApp/data/inMemoryStore"
	cli "toDoApp/pkg/cli/validation"
	"toDoApp/pkg/core"
)

func Get(args ...string) string {
	valid, errors := cli.ValidateGet(args)
	if !valid {
		return errors
	}

	userName := args[0]

	request := core.GetToDoRequest{UserName: userName}
	response := inMemoryStore.Fetch(request)
	return fmt.Sprintf("Entries: %q ", response)
}
