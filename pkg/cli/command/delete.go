package cliCommand

import (
	"fmt"
	"toDoApp/data/inMemoryStore"
	cliValidate "toDoApp/pkg/cli/validation"
	"toDoApp/pkg/core"
)

func Delete(args ...string) string {
	valid, errors := cliValidate.ValidateDelete(args)
	if !valid {
		return errors
	}

	userName := args[0]
	id := args[1]

	request := core.DeleteToDoRequest{UserName: userName, Id: id}
	pl("Delete request", request)
	response := inMemoryStore.Delete(request)
	return fmt.Sprintf("Deleted: %q", response)
}
