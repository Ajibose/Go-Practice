package handler

import (
	"html/template"
	"net/http"
)

func Home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "Page not found", http.StatusNotFound)
			return
		}

		tmpl, err := template.ParseFiles("./layout.html", "./home.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)

		tmpl.ExecuteTemplate(w, "layout", "")

	}
}
