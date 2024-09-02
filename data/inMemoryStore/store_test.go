package inMemoryStore

import (
	"reflect"
	"testing"
	"toDoApp/pkg/core"
)

func TestCreate(t *testing.T) {
	store = make(map[string]ToDoRecords)
	todos := []core.ToDo{{Id: "1", Instruction: "Do something"}, {Id: "2", Instruction: "Do something else"}}
	request := core.PostToDoRequest{UserName: "Dirk", ToDos: todos}
	expected := core.PostToDoResponse{ToDos: todos}

	actual := Create(request)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %q expected %q are not equal", actual, expected)
	}
}

func TestGet(t *testing.T) {
	store = make(map[string]ToDoRecords)
	todos := []core.ToDo{{Id: "1", Instruction: "Do something"}, {Id: "2", Instruction: "Do something else"}}
	request := core.PostToDoRequest{UserName: "Dirk", ToDos: todos}
	expected := core.GetToDoResponse{ToDos: todos}

	Create(request)
	actual := Fetch(core.GetToDoRequest{UserName: "Dirk"})

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %q expected %q are not equal", actual, expected)
	}
}

func TestUpdate(t *testing.T) {
	updateTestString := "Updated"
	currentTestString := "Current"
	store = make(map[string]ToDoRecords)
	todos := []core.ToDo{{Id: "1", Instruction: currentTestString}, {Id: "2", Instruction: "Do something else"}}
	request := core.PostToDoRequest{UserName: "Dirk", ToDos: todos}
	todoUpdate := core.ToDo{Id: "1", Instruction: updateTestString}
	expected := core.PutToDoResponse{ToDo: todoUpdate}

	Create(request)

	if store["Dirk"].Entries["1"] != currentTestString {
		t.Errorf("The todo should be update to '%s' but was  %s", currentTestString, store["Dirk"].Entries["1"])
	}

	actual := Update(core.PutToDoRequest{UserName: "Dirk", ToDo: todoUpdate})

	if store["Dirk"].Entries["1"] != updateTestString {
		t.Errorf("The todo should be update to '%s' but was  %s", updateTestString, store["Dirk"].Entries["1"])
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %q expected %q are not equal", actual, expected)
	}

}

func TestDelete(t *testing.T) {
	store = make(map[string]ToDoRecords)
	todos := []core.ToDo{{Id: "1", Instruction: "Do something"}, {Id: "2", Instruction: "Do something else"}}
	request := core.PostToDoRequest{UserName: "Dirk", ToDos: todos}
	expected := core.DeleteToDoResponse{ToDo: todos[0]}

	Create(request)

	if len(store["Dirk"].Entries) != 2 {
		t.Errorf("There should be 2 entires for Dirk but have %d", len(store["Dirk"].Entries))
	}

	actual := Delete(core.DeleteToDoRequest{UserName: "Dirk", Id: "1"})

	if len(store["Dirk"].Entries) != 1 {
		t.Errorf("There should be 1 entry for Dirk but have %d", len(store["Dirk"].Entries))
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual %q expected %q are not equal", actual, expected)
	}

}
