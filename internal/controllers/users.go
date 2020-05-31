package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/daichi5/web_dev_portal/internal/models"
)

var validPath = regexp.MustCompile("^/users/([0-9]+)$")

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if len(m) > 0 {
		switch r.Method {
		case "GET":
			show(w, r)
		}
	} else {
		switch r.Method {
		case "GET":
			index(w, r)
		case "POST":
			create(w, r)
		}
	}
}

func show(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(models.User{Name: r.URL.Path})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func index(w http.ResponseWriter, r *http.Request) {
	users := []models.User{}
	models.Db.Find(&users)
	json, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func create(w http.ResponseWriter, r *http.Request) {
	user := models.User{Name: r.FormValue("name")}
	models.Db.Create(&user)
	json, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
