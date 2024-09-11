package core

type ToDo struct {
	Id          string
	Instruction string
}

type GetToDoRequest struct {
	UserName string
}

type GetToDoResponse struct {
	ToDos []ToDo
}

type PostToDoRequest struct {
	UserName string
	ToDos    []string
}

type PostToDoResponse struct {
	ToDos []ToDo
}

type PutToDoRequest struct {
	UserName string
	ToDo     ToDo
}

type PutToDoResponse struct {
	ToDos []ToDo
}

type DeleteToDoRequest struct {
	UserName string
	Id       string
}

type DeleteToDoResponse struct {
	ToDos []ToDo
}

type ToDoRequest interface {
	GetToDoRequest | PostToDoRequest | PutToDoRequest | DeleteToDoRequest
}
