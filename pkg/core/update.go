package core

import (
	"fmt"
	"strings"
	"toDoApp/data/inMemoryStore"
	coreTypes "toDoApp/pkg/types"
)

func Update(args ...string) string {
	if len(args) != 2 {
		return "Incorrect number of arguments type 'help' for more information"
	}

	userName := args[0]
	todo := strings.Split(args[1], ":")

	request := coreTypes.PutToDoRequest{
		UserName: userName,
		ToDo:     coreTypes.ToDo{Id: todo[0], Instruction: todo[1]}}
	response := inMemoryStore.Update(request)
	return fmt.Sprintln("Updated: ", response)
}
