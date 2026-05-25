package main

import (
	"net/http"
	"practice/http/handler"
	"practice/http/model"
)

func main() {
	mux := http.NewServeMux()
	store := &model.Storage{Users: []model.User{}, Ids: make(map[string]model.Data)}

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	mux.HandleFunc("POST /paste", handler.CreatePaste(store))
	mux.HandleFunc("GET /paste", handler.GetAllPaste(store))
	mux.HandleFunc("GET /paste/{id}", handler.GetPaste(store))

	mux.HandleFunc("/", handler.Home())

	mux.HandleFunc("GET /paste/{id}/info", handler.GetPasteInfo(store))

	mux.HandleFunc("GET /paste/{id}/view", handler.ViewPaste(store))

	mux.HandleFunc("POST /users", handler.CreateUser(store))

	http.ListenAndServe(":8080", logger(mux))

}
