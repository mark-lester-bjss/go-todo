package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"toDoApp/pkg/cli"
	cliCommand "toDoApp/pkg/cli/command"
)

var pl = fmt.Println

func init() {
	cli.RegisterCommand("add", cliCommand.Create)
	cli.RegisterCommand("update", cliCommand.Update)
	cli.RegisterCommand("get", cliCommand.Get)
	cli.RegisterCommand("delete", cliCommand.Delete)
	cli.RegisterCommand("help", cliCommand.Help)
}

func main() {
	var cliParams string

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("todo app: ")
		cliParams, _ = reader.ReadString('\n')
		uiMessage := cli.ExecuteCommand(removeNewLine(cliParams))
		pl(uiMessage)
	}
}

func removeNewLine(s string) string {
	return strings.Replace(s, "\n", "", -1)
}
