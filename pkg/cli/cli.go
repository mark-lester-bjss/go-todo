package cli

import (
	"fmt"
	"strings"
)

var commands = make(map[string]func(str ...string) string)

func RegisterCommand(name string, function func(str ...string) string) {
	commands[name] = function
}

func ExecuteCommand(cliRequest string) string {
	commandName, args := separtateParams(cliRequest)
	command := commands[commandName]
	if command == nil {
		return fmt.Sprintf("The command '%s' is not a valid command. Type help for more information", commandName)
	}
	return command(args...)
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
