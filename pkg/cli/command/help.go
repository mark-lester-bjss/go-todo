package cliCommand

import "toDoApp/data/database"

func Help(database database.ToDoDataStore, args ...string) string {
	return "A detailed help message: "
}
