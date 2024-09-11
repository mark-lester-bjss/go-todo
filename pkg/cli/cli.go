package cli

import (
	"fmt"
	"strings"
	"toDoApp/data/database"
)

var commands = make(map[string]func(database database.ToDoDataStore, str ...string) string)

func RegisterCommand(name string, function func(database database.ToDoDataStore, str ...string) string) {
	commands[name] = function
}

func ExecuteCommand(database database.ToDoDataStore, cliRequest string) string {
	commandName, args := separtateParams(cliRequest)
	command := commands[commandName]
	if command == nil {
		return fmt.Sprintf("The command '%s' is not a valid command. Type help for more information", commandName)
	}
	return command(database, args...)
}

func separtateParams(cliParams string) (commandName string, params []string) {
	splitParams := strings.Fields(cliParams)
	commandName = splitParams[0]
	userName := splitParams[1]

	cliParams, _ = strings.CutPrefix(cliParams, fmt.Sprint(commandName, " "))
	cliParams, _ = strings.CutPrefix(cliParams, fmt.Sprint(userName, " "))

	params = make([]string, 0)
	params = append(params, userName)
	params = append(params, cliParams)

	return commandName, params
}
