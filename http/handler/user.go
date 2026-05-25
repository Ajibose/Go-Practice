package handler

import (
	"encoding/json"
	"net/http"
	"practice/http/model"
)


func CreateUser(store *model.Storage) http.HandlerFunc {
	users := store.Users
	return func(w http.ResponseWriter, r *http.Request) {
		var u model.User

		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, "Bad json provided", http.StatusBadRequest)
			return
		}

		if u.Name == "" {
			http.Error(w, "Name not allowed empty", http.StatusBadRequest)
			return
		}

		if u.Email == "" {
			http.Error(w, "Name not allowed empty", http.StatusBadRequest)
			return
		}

		users = append(users, u)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)

		json.NewEncoder(w).Encode(u)
	}
}
