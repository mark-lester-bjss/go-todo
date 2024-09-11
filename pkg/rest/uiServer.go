package rest

import (
	"fmt"
	"html/template"
	"net/http"
	"toDoApp/data/database"
	"toDoApp/pkg/core"
)

var pl = fmt.Println

type Profile struct {
	FirstName string
	LastName  string
}

func UiServer(database database.ToDoDataStore) {
	pl("Starting Ui server")
	mux := http.NewServeMux()

	mux.HandleFunc("/home", func(writer http.ResponseWriter, request *http.Request) {
		template, err := template.ParseFiles("pkg/rest/home.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		template.Execute(writer, "")
	})

	mux.HandleFunc("/interact", func(writer http.ResponseWriter, request *http.Request) {
		template, err := template.ParseFiles("pkg/rest/interact.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		template.Execute(writer, "")
	})

	mux.HandleFunc("/fetch", func(writer http.ResponseWriter, httpRequest *http.Request) {
		template, err := template.ParseFiles("pkg/rest/home.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		request := core.GetToDoRequest{UserName: httpRequest.FormValue("userName")}
		valid, errors := core.ValidateGetToDoRequest(request)
		if valid {
			err := template.Execute(writer, database.Fetch(request))
			if err != nil {
				http.Error(writer, err.Error(), http.StatusUnauthorized)
			}
		} else {
			fmt.Fprint(writer, errors)
		}
	})

	mux.HandleFunc("/create", func(writer http.ResponseWriter, httpRequest *http.Request) {
		template, err := template.ParseFiles("pkg/rest/interact.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		request := core.PostToDoRequest{UserName: httpRequest.FormValue("userName"), ToDos: []string{httpRequest.FormValue("instruction")}}
		valid, validationErrors := core.ValidatePostToDoRequest(request)
		if valid {
			err := template.Execute(writer, database.Create(request))
			if err != nil {
				http.Error(writer, err.Error(), http.StatusUnauthorized)
			}
		} else {
			fmt.Fprint(writer, validationErrors)
		}
	})

	mux.HandleFunc("/update", func(writer http.ResponseWriter, httpRequest *http.Request) {
		template, err := template.ParseFiles("pkg/rest/interact.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		request := core.PutToDoRequest{UserName: httpRequest.FormValue("userName"), ToDo: core.ToDo{Id: httpRequest.FormValue("id"), Instruction: httpRequest.FormValue("instruction")}}
		valid, validationErrors := core.ValidatePutToDoRequest(request)
		if valid {
			err := template.Execute(writer, database.Update(request))
			if err != nil {
				http.Error(writer, err.Error(), http.StatusUnauthorized)
			}
		} else {
			fmt.Fprint(writer, validationErrors)
		}
	})

	mux.HandleFunc("/delete", func(writer http.ResponseWriter, httpRequest *http.Request) {
		template, err := template.ParseFiles("pkg/rest/interact.html")
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}

		request := core.DeleteToDoRequest{UserName: httpRequest.FormValue("userName"), Id: httpRequest.FormValue("id")}
		valid, validationErrors := core.ValidateDeleteToDoRequest(request)
		if valid {
			err := template.Execute(writer, database.Delete(request))
			if err != nil {
				http.Error(writer, err.Error(), http.StatusUnauthorized)
			}
		} else {
			fmt.Fprint(writer, validationErrors)
		}
	})

	mux.HandleFunc("GET /help", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "help")
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
	pl("Ui server is running")
}
