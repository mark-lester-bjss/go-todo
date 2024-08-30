package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	cli "toDoApp/pkg/cli"
	core "toDoApp/pkg/core"
)

var pl = fmt.Println

func init() {
	cli.RegisterCommand("add", core.Create)
	cli.RegisterCommand("update", core.Update)
	cli.RegisterCommand("get", core.Get)
	cli.RegisterCommand("delete", core.Delete)
	cli.RegisterCommand("help", core.Help)
}

func main() {
	var cliParams string

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("todo app: ")
		cliParams, _ = reader.ReadString('\n')
		uiMessage := cli.ExecuteCommand(removeNewLine(cliParams))
		pl("Please read: ", uiMessage)
	}
}

func removeNewLine(s string) string {
	return strings.Replace(s, "\n", "", -1)
}
