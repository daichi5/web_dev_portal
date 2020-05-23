package main

import (
	"net/http"
	"text/template"
)

type PageData struct {
	Message string
}

var templates = template.Must(template.ParseFiles("internal/views/home.html"))

func viewHomeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.html", PageData{Message: "this is message"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", viewHomeHandler)
	http.ListenAndServe(":8080", nil)
}
