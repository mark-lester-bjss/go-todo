package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"toDoApp/data/database"
	"toDoApp/data/inMemoryStore"
	"toDoApp/pkg/cli"
	cliCommand "toDoApp/pkg/cli/command"
	"toDoApp/pkg/rest"
)

var pl = fmt.Println
var db database.ToDoDataStore

func main() {
	db := flag.String("db", "memory", "Which database to use")
	ui := flag.String("ui", "cli", "Which ui to use")
	flag.Parse()

	pl("main called db: ", *db)
	pl("main called ui: ", *ui)

	handleDbFlag(*db)
	handleUiFlag(*ui)
}

func removeNewLine(s string) string {
	return strings.Replace(s, "\n", "", -1)
}

func handleDbFlag(flag string) {
	if flag == "memory" {
		db = inMemoryStore.InMemoryStore{}
	}
}

func handleUiFlag(flag string) {
	if flag == "cli" {
		cli.RegisterCommand("add", cliCommand.Create)
		cli.RegisterCommand("update", cliCommand.Update)
		cli.RegisterCommand("get", cliCommand.Get)
		cli.RegisterCommand("delete", cliCommand.Delete)
		cli.RegisterCommand("help", cliCommand.Help)

		var cliParams string
		reader := bufio.NewReader(os.Stdin)
		for {
			fmt.Print("todo app: ")
			cliParams, _ = reader.ReadString('\n')
			uiMessage := cli.ExecuteCommand(db, removeNewLine(cliParams))
			pl(uiMessage)
		}
	}
	if flag == "rest" {
		rest.RestServer(db)
	}

	if flag == "ui" {
		rest.UiServer(db)
	}

}
