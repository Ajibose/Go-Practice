package main

import (
	"html/template"
	"log"
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

	// mux.HandleFunc("POST /users", handler.CreateUser(store))
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl, err := template.ParseFiles("./templates/layout.html", "./templates/register.html")
			if err != nil {
				log.Println("Error pasrsing template", err)
				http.Error(w, "Error Parsing templater", 500)
				return
			}

			w.Header().Set("Content-Type", "text/html")
			tmpl.ExecuteTemplate(w, "layout", nil)
			return
		}

		handler.CreateUser(store).ServeHTTP(w, r)
	})

	log.Println("Listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", logger(mux))
	log.Fatal(err)

}
