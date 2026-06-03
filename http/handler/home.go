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

		tmpl, err := template.ParseFiles("./templates/layout.html", "./templates/home.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")

		tmpl.ExecuteTemplate(w, "layout", "")

	}
}
