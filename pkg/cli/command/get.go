package cliCommand

import (
	"fmt"
	"toDoApp/data/database"
	cli "toDoApp/pkg/cli/validation"
	"toDoApp/pkg/core"
)

func Get(database database.ToDoDataStore, args ...string) string {
	valid, errors := cli.ValidateGet(args)
	if !valid {
		return errors
	}

	userName := args[0]

	request := core.GetToDoRequest{UserName: userName}
	response := database.Fetch(request)
	return fmt.Sprintf("Entries: %q ", response)
}
