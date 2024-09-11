package inMemoryStore

import (
	"reflect"
	"testing"
	"toDoApp/pkg/core"
)

func TestCreate(t *testing.T) {
	database := InMemoryStore{}
	store = make(map[string]ToDoRecords)
	todos := []string{"Do something", "Do something else"}
	request := core.PostToDoRequest{UserName: "Dirk", ToDos: todos}
	expected := 2

	actual := database.Create(request)

	if len(actual.ToDos) != expected {
		t.Errorf("actual %q expected %q are not equal", actual, expected)
	}
}

func TestGet(t *testing.T) {
	database := InMemoryStore{}
	store = make(map[string]ToDoRecords)
	todos := []string{"Do something", "Do something else"}
	request := core.PostToDoRequest{UserName: "Dirk", ToDos: todos}
	expected := 2

	database.Create(request)
	actual := database.Fetch(core.GetToDoRequest{UserName: "Dirk"})

	if len(actual.ToDos) != expected {
		t.Errorf("actual %q expected %q are not equal", actual, expected)
	}
}

func TestUpdate(t *testing.T) {
	database := InMemoryStore{}
	updateTestString := "Updated"
	currentTestString := "Current"
	store = make(map[string]ToDoRecords)
	todos := []string{currentTestString}
	request := core.PostToDoRequest{UserName: "Dirk", ToDos: todos}

	database.Create(request)
	fetchRequest := core.GetToDoRequest{UserName: "Dirk"}
	fetchResponse := database.Fetch(fetchRequest)
	todoUpdate := core.ToDo{Id: fetchResponse.ToDos[0].Id, Instruction: updateTestString}
	expected := core.PutToDoResponse{ToDos: []core.ToDo{todoUpdate}}

	if store["Dirk"].Entries[fetchResponse.ToDos[0].Id] != currentTestString {
		t.Errorf("1The todo should be '%s' but was  %s", currentTestString, store["Dirk"].Entries[fetchResponse.ToDos[0].Id])
	}

	actual := database.Update(core.PutToDoRequest{UserName: "Dirk", ToDo: todoUpdate})

	if store["Dirk"].Entries[fetchResponse.ToDos[0].Id] != updateTestString {
		t.Errorf("2The todo should be updated to '%s' but was '%s'", updateTestString, store["Dirk"].Entries[fetchResponse.ToDos[0].Id])
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("3actual %q expected %q are not equal", actual, expected)
	}

}

func TestDelete(t *testing.T) {
	database := InMemoryStore{}
	store = make(map[string]ToDoRecords)
	todos := []string{"Do something", "Do something else"}
	request := core.PostToDoRequest{UserName: "Dirk", ToDos: todos}

	database.Create(request)
	fetchRequest := core.GetToDoRequest{UserName: "Dirk"}
	fetchResponse := database.Fetch(fetchRequest)
	expected := core.DeleteToDoResponse{ToDos: []core.ToDo{fetchResponse.ToDos[0]}}

	if len(store["Dirk"].Entries) != 2 {
		t.Errorf("There should be 2 entires for Dirk but have %d", len(store["Dirk"].Entries))
	}

	actual := database.Delete(core.DeleteToDoRequest{UserName: "Dirk", Id: fetchResponse.ToDos[0].Id})

	if len(store["Dirk"].Entries) != 1 {
		t.Errorf("There should be 1 entry for Dirk but have %d", len(store["Dirk"].Entries))
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %q expected %q are not equal", actual, expected)
	}

}
