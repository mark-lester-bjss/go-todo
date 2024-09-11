package cliCommand

import (
	"fmt"
	"strings"
	"toDoApp/data/database"
	cliValidate "toDoApp/pkg/cli/validation"
	"toDoApp/pkg/core"
)

func Create(database database.ToDoDataStore, args ...string) string {
	valid, errors := cliValidate.ValidateCreate(args)
	if !valid {
		return errors
	}

	userName := args[0]

	request := core.PostToDoRequest{UserName: userName, ToDos: strings.Split(args[1], ",")}
	response := database.Create(request)
	return fmt.Sprintf("Added: %q ", response)
}
