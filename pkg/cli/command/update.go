package cliCommand

import (
	"fmt"
	"strings"
	"toDoApp/data/database"
	cli "toDoApp/pkg/cli/validation"
	"toDoApp/pkg/core"
)

var pl = fmt.Println

func Update(database database.ToDoDataStore, args ...string) string {

	valid, errors := cli.ValidateUpdate(args)
	if !valid {
		return errors
	}

	userName := args[0]
	todo := strings.Split(args[1], ":")

	request := core.PutToDoRequest{
		UserName: userName,
		ToDo:     core.ToDo{Id: todo[0], Instruction: todo[1]}}
	pl("Update request", request)
	response := database.Update(request)
	return fmt.Sprintln("Updated: ", response)
}
