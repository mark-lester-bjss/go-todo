package cliCommand

import (
	"fmt"
	"strings"
	"toDoApp/data/inMemoryStore"
	cliValidate "toDoApp/pkg/cli/validation"
	"toDoApp/pkg/core"
)

func Create(args ...string) string {
	valid, errors := cliValidate.ValidateCreate(args)
	if !valid {
		return errors
	}

	userName := args[0]

	request := core.PostToDoRequest{UserName: userName, ToDos: strings.Split(args[1], ",")}
	response := inMemoryStore.Create(request)
	return fmt.Sprintf("Added: %q ", response)
}
