package main

import (
	"net/http"
	"text/template"

	"github.com/daichi5/web_dev_portal/internal/controllers"
)

type PageData struct {
	Message string
}

var templates = template.Must(template.ParseFiles("internal/views/home.html"))

func viewHomeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.html", PageData{Message: "this is a message"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/users/", controllers.UsersHandler)
	http.HandleFunc("/", viewHomeHandler)
	http.ListenAndServe(":8080", nil)
}
