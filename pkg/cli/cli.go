package cli

import (
	"strings"
)

var commands = make(map[string]func(str ...string) string)

func RegisterCommand(name string, function func(str ...string) string) {
	commands[name] = function
}

func ExecuteCommand(cliRequest string) string {
	commandName, args := separtateParams(cliRequest)
	command := commands[commandName]
	return command(args...)
}

func separtateParams(cliParams string) (commandName string, params []string) {
	splitParams := strings.Fields(cliParams)
	commandName = splitParams[0]
	params = make([]string, 0)

	for i := 1; i < len(splitParams); i++ {
		params = append(params, splitParams[i])
	}

	return commandName, params

}
