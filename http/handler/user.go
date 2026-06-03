package handler

import (
	"encoding/json"
	"net/http"
	"practice/http/model"
)

func CreateUser(store *model.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User

		if r.Header.Get("Content-Type") == "application/json" {
			err := json.NewDecoder(r.Body).Decode(&u)
			if err != nil {
				http.Error(w, "Bad json provided", http.StatusBadRequest)
				return
			}
		} else {
			u.Name = r.FormValue("name")
			u.Email = r.FormValue("email")
		}

		if u.Name == "" {
			http.Error(w, "Name not allowed empty", http.StatusBadRequest)
			return
		}

		if u.Email == "" {
			http.Error(w, "Name not allowed empty", http.StatusBadRequest)
			return
		}

		store.Users = append(store.Users, u)

		if r.Header.Get("Content-Type") == "application/json" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(201)
			json.NewEncoder(w).Encode(u)
			return
		}
		http.Redirect(w, r, "/", 303)
	}
}
