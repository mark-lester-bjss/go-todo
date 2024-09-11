package database

import "toDoApp/pkg/core"

type ToDoDataStore interface {
	Fetch(request core.GetToDoRequest) (response core.GetToDoResponse)
	Create(request core.PostToDoRequest) (response core.PostToDoResponse)
	Update(request core.PutToDoRequest) (response core.PutToDoResponse)
	Delete(request core.DeleteToDoRequest) (response core.DeleteToDoResponse)
}
