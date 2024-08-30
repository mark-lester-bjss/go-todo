package core

import (
	"fmt"
	"toDoApp/data/inMemoryStore"
	coreTypes "toDoApp/pkg/types"
)

func Get(args ...string) string {
	if len(args) != 1 {
		return "Incorrect number of arguments type 'help' for more information"
	}

	userName := args[0]

	request := coreTypes.GetToDoRequest{UserName: userName}
	response := inMemoryStore.Fetch(request)
	return fmt.Sprintf("Entries: %q ", response)
}
