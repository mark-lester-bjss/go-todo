package cliCommand

import (
	"fmt"
	"strings"
	"toDoApp/data/inMemoryStore"
	cliValidate "toDoApp/pkg/cli/validation"
	"toDoApp/pkg/core"

	"github.com/google/uuid"
)

func Create(args ...string) string {
	valid, errors := cliValidate.ValidateCreate(args)
	if !valid {
		return errors
	}

	userName := args[0]
	todos := createTodosArray(strings.Split(args[1], ","))

	request := core.PostToDoRequest{UserName: userName, ToDos: todos}
	response := inMemoryStore.Create(request)
	return fmt.Sprintf("Added: %q ", response)
}

func createTodosArray(strings []string) []core.ToDo {
	todos := make([]core.ToDo, 0)
	for i := 0; i < len(strings); i++ {
		todo := core.ToDo{Id: uuid.New().String(), Instruction: strings[i]}
		todos = append(todos, todo)
	}
	return todos
}
