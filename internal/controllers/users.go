package controllers

import (
	"encoding/json"
	"net/http"
	"regexp"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var validPath = regexp.MustCompile("^/users/([0-9]+)$")

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	m := validPath.FindStringSubmatch(r.URL.Path)
	if len(m) > 0 {
		switch r.Method {
		case "GET":
			show(w, r)
		}
	} else {
		index(w, r)
	}

}

func show(w http.ResponseWriter, r *http.Request) {
	json, err := json.Marshal(User{ID: 1, Name: r.URL.Path})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func index(w http.ResponseWriter, r *http.Request) {
	users := []User{}
	users = append(users, User{ID: 1, Name: "tom"})
	users = append(users, User{ID: 2, Name: "alice"})
	json, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
