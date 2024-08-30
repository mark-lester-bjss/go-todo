package core

import (
	"fmt"
	"toDoApp/data/inMemoryStore"
	coreTypes "toDoApp/pkg/types"
)

func Delete(args ...string) string {
	if len(args) != 2 {
		return "Incorrect number of arguments type 'help' for more information"
	}

	userName := args[0]
	id := args[1]

	request := coreTypes.DeleteToDoRequest{UserName: userName, Id: id}
	response := inMemoryStore.Delete(request)
	return fmt.Sprintf("Deleted: %q", response)
}
