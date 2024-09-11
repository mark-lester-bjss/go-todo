package inMemoryStore

import (
	"fmt"
	"toDoApp/pkg/core"

	"github.com/google/uuid"
)

var pl = fmt.Println

type InMemoryStore struct{}

var store = make(map[string]ToDoRecords)

func (i InMemoryStore) Fetch(request core.GetToDoRequest) (response core.GetToDoResponse) {
	record, ok := store[request.UserName]

	if ok {
		for key, value := range record.Entries {
			response.ToDos = append(response.ToDos, core.ToDo{Id: key, Instruction: value})
		}
	}
	return response
}

func (i InMemoryStore) Create(request core.PostToDoRequest) (response core.PostToDoResponse) {
	record, ok := store[request.UserName]

	if ok {
		response = addToExistingRecord(request, record)
	} else {
		response = createNewRecord(request)
	}
	return response
}

func createTodosArray(strings []string) (todos []core.ToDo) {
	for i := 0; i < len(strings); i++ {
		todo := core.ToDo{Id: uuid.New().String(), Instruction: strings[i]}
		todos = append(todos, todo)
	}
	return todos
}

func addToExistingRecord(request core.PostToDoRequest, record ToDoRecords) core.PostToDoResponse {
	todoArray := createTodosArray(request.ToDos)
	for _, todo := range todoArray {
		record.Entries[todo.Id] = todo.Instruction
	}
	return core.PostToDoResponse{ToDos: todoArray}
}

func createNewRecord(request core.PostToDoRequest) core.PostToDoResponse {
	todoArray := createTodosArray(request.ToDos)
	newRecord := ToDoRecords{make(map[string]string)}
	for _, todo := range todoArray {
		newRecord.Entries[todo.Id] = todo.Instruction
	}
	store[request.UserName] = newRecord
	return core.PostToDoResponse{ToDos: todoArray}
}

func (i InMemoryStore) Update(request core.PutToDoRequest) (response core.PutToDoResponse) {
	record, ok := store[request.UserName]

	if ok {
		_, ok := record.Entries[request.ToDo.Id]
		if ok {
			record.Entries[request.ToDo.Id] = request.ToDo.Instruction
			response = core.PutToDoResponse{ToDos: []core.ToDo{request.ToDo}}
		}
	}
	return response
}

func (i InMemoryStore) Delete(request core.DeleteToDoRequest) (response core.DeleteToDoResponse) {
	record, ok := store[request.UserName]

	if ok {
		entry, ok := record.Entries[request.Id]
		if ok {
			response = core.DeleteToDoResponse{ToDos: []core.ToDo{{Id: request.Id, Instruction: entry}}}
			delete(record.Entries, request.Id)
		}
	}
	return response
}
