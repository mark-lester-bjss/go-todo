package rest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"toDoApp/data/database"
	"toDoApp/pkg/core"
)

func RestServer(database database.ToDoDataStore) {
	pl("Starting REST serve")
	mux := http.NewServeMux()

	mux.HandleFunc("GET /todo/{userName}", func(writer http.ResponseWriter, request *http.Request) {
		getToDoRequest := core.GetToDoRequest{UserName: request.PathValue("userName")}
		valid, errors := core.ValidateGetToDoRequest(getToDoRequest)
		if valid {
			fmt.Fprint(writer, database.Fetch(getToDoRequest))
		} else {
			fmt.Fprint(writer, errors)
		}
	})

	mux.HandleFunc("POST /todo", func(writer http.ResponseWriter, httpRequest *http.Request) {
		putToDoRequest := core.PostToDoRequest{}
		success, errors, request := convertToRequestType(putToDoRequest, httpRequest, core.ValidatePostToDoRequest)
		if success {
			fmt.Fprint(writer, database.Create(request))
		} else {
			fmt.Fprint(writer, errors)
		}
	})

	mux.HandleFunc("PUT /todo", func(writer http.ResponseWriter, httpRequest *http.Request) {
		putToDoRequest := core.PutToDoRequest{}
		success, errors, request := convertToRequestType(putToDoRequest, httpRequest, core.ValidatePutToDoRequest)
		if success {
			fmt.Fprint(writer, database.Update(request))
		} else {
			fmt.Fprint(writer, errors)
		}
	})

	mux.HandleFunc("DELETE /todo", func(writer http.ResponseWriter, httpRequest *http.Request) {
		deleteToDoRequest := core.DeleteToDoRequest{}
		success, errors, request := convertToRequestType(deleteToDoRequest, httpRequest, core.ValidateDeleteToDoRequest)
		if success {
			fmt.Fprint(writer, database.Delete(request))
		} else {
			fmt.Fprint(writer, errors)
		}
	})

	mux.HandleFunc("GET /help", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "help")
	})
	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
	pl("REST server is running")
}

func convertToRequestType[R core.ToDoRequest](requestIn R, httpRequest *http.Request, validate func(R) (valid bool, errors []string)) (success bool, errors []string, requestOut R) {
	errors = []string{}
	body, readError := io.ReadAll(httpRequest.Body)
	if readError != nil {
		return false, []string{"Failed to read request body"}, requestIn
	}

	marshalError := json.Unmarshal(body, &requestIn)
	if marshalError != nil {
		return false, []string{"Failed to transform request"}, requestIn
	}

	valid, validationErrors := validate(requestIn)
	if !valid {
		errors = validationErrors
	}
	return len(errors) == 0, errors, requestIn
}
