package core

func ValidateToDo(request ToDo) (valid bool, errors []string) {
	if request.Id == "" {
		errors = append(errors, "A valid uuid must be provided for the todo Id")
	}
	if request.Instruction == "" {
		errors = append(errors, "A new todo instruction must be provided for the todo")
	}
	return len(errors) == 0, errors
}

func ValidateGetToDoRequest(request GetToDoRequest) (valid bool, errors []string) {
	if request.UserName == "" {
		errors = append(errors, "A user name must be provided")
	}
	return len(errors) == 0, errors
}

func ValidatePostToDoRequest(request PostToDoRequest) (valid bool, errors []string) {
	if request.UserName == "" {
		errors = append(errors, "A user name must be provided")
	}
	if len(request.ToDos) == 0 {
		errors = append(errors, "At least one todo must be provided")
	}

	return len(errors) == 0, errors
}

func ValidatePutToDoRequest(request PutToDoRequest) (valid bool, errors []string) {
	if request.UserName == "" {
		errors = append(errors, "A user name must be provided")
	}
	valid, validationErrors := ValidateToDo(request.ToDo)
	if !valid {
		errors = append(errors, validationErrors...)
	}
	return len(errors) == 0, errors
}

func ValidateDeleteToDoRequest(request DeleteToDoRequest) (valid bool, errors []string) {
	if request.UserName == "" {
		errors = append(errors, "A user name must be provided")
	}
	if request.Id == "" {
		errors = append(errors, "A valid uuid must be provided for the todo Id")
	}
	return len(errors) == 0, errors
}
