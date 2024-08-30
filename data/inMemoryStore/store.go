package inMemoryStore

import (
	coreTypes "toDoApp/pkg/types"
)

var store = make(map[string]ToDoRecords)

func Fetch(request coreTypes.GetToDoRequest) coreTypes.GetToDoResponse {
	response := coreTypes.GetToDoResponse{}
	record, ok := store[request.UserName]

	if ok {
		for key, value := range record.Entries {
			response.ToDos = append(response.ToDos, coreTypes.ToDo{Id: key, Instruction: value})
		}
	}
	return response
}

func Create(request coreTypes.PostToDoRequest) coreTypes.PostToDoResponse {
	response := coreTypes.PostToDoResponse{}
	record, ok := store[request.UserName]

	if ok {
		response = addToExistingRecord(request, record)
	} else {
		response = createNewRecord(request)
	}
	return response
}

func addToExistingRecord(request coreTypes.PostToDoRequest, record ToDoRecords) coreTypes.PostToDoResponse {
	for _, todo := range request.ToDos {
		record.Entries[todo.Id] = todo.Instruction
	}
	return coreTypes.PostToDoResponse{ToDos: request.ToDos}
}

func createNewRecord(request coreTypes.PostToDoRequest) coreTypes.PostToDoResponse {
	newRecord := ToDoRecords{make(map[string]string)}
	for _, todo := range request.ToDos {
		newRecord.Entries[todo.Id] = todo.Instruction
	}
	store[request.UserName] = newRecord
	return coreTypes.PostToDoResponse{ToDos: request.ToDos}
}

func Update(request coreTypes.PutToDoRequest) coreTypes.PutToDoResponse {
	response := coreTypes.PutToDoResponse{}
	record, ok := store[request.UserName]

	if ok {
		_, ok := record.Entries[request.ToDo.Id]
		if ok {
			record.Entries[request.ToDo.Id] = request.ToDo.Instruction
			response = coreTypes.PutToDoResponse{ToDo: request.ToDo}
		}
	}
	return response
}

func Delete(request coreTypes.DeleteToDoRequest) coreTypes.DeleteToDoResponse {
	response := coreTypes.DeleteToDoResponse{}
	record, ok := store[request.UserName]

	if ok {
		entry, ok := record.Entries[request.Id]
		if ok {
			response = coreTypes.DeleteToDoResponse{ToDo: coreTypes.ToDo{Id: request.Id, Instruction: entry}}
			delete(record.Entries, request.Id)
		}
	}
	return response
}
