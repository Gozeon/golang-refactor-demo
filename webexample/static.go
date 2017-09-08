package main

import (
	"net/http"
)

func main()  {
	fs:=http.FileServer(http.Dir("./webexample/assets/"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.ListenAndServe("localhost:8080", nil)
}
