package core

import (
	"fmt"
	"strings"
	"toDoApp/data/inMemoryStore"
	coreTypes "toDoApp/pkg/types"

	"github.com/google/uuid"
)

func Create(args ...string) string {
	if len(args) != 2 {
		return "Incorrect number of arguments type 'help' for more information"
	}

	userName := args[0]
	todos := createTodosArray(strings.Split(args[1], ","))

	request := coreTypes.PostToDoRequest{UserName: userName, ToDos: todos}
	response := inMemoryStore.Create(request)
	return fmt.Sprintf("Added: %q ", response)
}

func createTodosArray(strings []string) []coreTypes.ToDo {
	todos := make([]coreTypes.ToDo, 0)
	for i := 0; i < len(strings); i++ {
		todo := coreTypes.ToDo{Id: uuid.New().String(), Instruction: strings[i]}
		todos = append(todos, todo)
	}
	return todos
}
