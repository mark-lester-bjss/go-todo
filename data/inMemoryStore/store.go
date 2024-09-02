package inMemoryStore

import (
	"fmt"
	"toDoApp/pkg/core"
)

var pl = fmt.Println

var store = make(map[string]ToDoRecords)

func Fetch(request core.GetToDoRequest) core.GetToDoResponse {
	response := core.GetToDoResponse{}
	record, ok := store[request.UserName]

	if ok {
		for key, value := range record.Entries {
			response.ToDos = append(response.ToDos, core.ToDo{Id: key, Instruction: value})
		}
	}
	return response
}

func Create(request core.PostToDoRequest) core.PostToDoResponse {
	response := core.PostToDoResponse{}
	record, ok := store[request.UserName]

	if ok {
		response = addToExistingRecord(request, record)
	} else {
		response = createNewRecord(request)
	}
	return response
}

func addToExistingRecord(request core.PostToDoRequest, record ToDoRecords) core.PostToDoResponse {
	for _, todo := range request.ToDos {
		record.Entries[todo.Id] = todo.Instruction
	}
	return core.PostToDoResponse{ToDos: request.ToDos}
}

func createNewRecord(request core.PostToDoRequest) core.PostToDoResponse {
	newRecord := ToDoRecords{make(map[string]string)}
	for _, todo := range request.ToDos {
		newRecord.Entries[todo.Id] = todo.Instruction
	}
	store[request.UserName] = newRecord
	return core.PostToDoResponse{ToDos: request.ToDos}
}

func Update(request core.PutToDoRequest) core.PutToDoResponse {
	pl("Update request", request)
	response := core.PutToDoResponse{}
	record, ok := store[request.UserName]

	if ok {
		_, ok := record.Entries[request.ToDo.Id]
		if ok {
			record.Entries[request.ToDo.Id] = request.ToDo.Instruction
			response = core.PutToDoResponse{ToDo: request.ToDo}
		}
	}
	return response
}

func Delete(request core.DeleteToDoRequest) core.DeleteToDoResponse {
	pl("Delete request", request)
	response := core.DeleteToDoResponse{}
	record, ok := store[request.UserName]

	if ok {
		entry, ok := record.Entries[request.Id]
		if ok {
			response = core.DeleteToDoResponse{ToDo: core.ToDo{Id: request.Id, Instruction: entry}}
			delete(record.Entries, request.Id)
		}
	}
	return response
}
