package main

import (
	"html/template"
	"net/http"
	"path/filepath"
)

type Todo struct {
	Title string
	Done  bool
}

type todoPageData struct {
	PageTitle string
	Todos     []Todo
}

func main() {
	filePrefix, _ := filepath.Abs("./webexample/")
	tmpl := template.Must(template.ParseFiles(filePrefix+ "/template.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := todoPageData{
			PageTitle: "TODO List",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe("localhost:8080", nil)
}
