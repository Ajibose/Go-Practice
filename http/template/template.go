package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, r *http.Request) {
		data := struct {
			Name       string
			PasteCount int
		}{
			"Ibrahim",
			23,
		}

		tmpl, err := template.ParseFiles("./index.html")
		if err != nil {
			fmt.Println("Error parse html")
			http.Error(w, "Error parsing file", http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8080", mux)

}
