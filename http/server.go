package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"math/big"
	"net/http"
	"strconv"
	"time"
)

type Data struct {
	ID        string
	CreatedAt string
	Body      string
}

var ids = make(map[string]Data)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /paste", func(w http.ResponseWriter, r *http.Request) {

		randomNum, _ := rand.Int(rand.Reader, big.NewInt(9000))
		id := randomNum.Int64() + 1000
		idString := strconv.Itoa(int(id))

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading body", err)
			http.Error(w, "failed to read body", http.StatusInternalServerError)
			return
		}

		d := Data{
			idString,
			time.Now().Format("2006-01-02 15:04:05"),
			string(body),
		}

		if len(body) == 0 {
			http.Error(w, "Body not provided", 400)
			return
		}

		ids[idString] = d
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(201)
		fmt.Fprintf(w, strconv.Itoa(int(id)))
	})

	mux.HandleFunc("GET /paste", func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./layout.html", "./paste.html")
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error parsing templtate", 500)
			return
		}

		tmpl.ExecuteTemplate(w, "layout", ids)
	})

	mux.HandleFunc("GET /paste/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		d, exist := ids[id]

		if !exist {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
		fmt.Fprintf(w, d.Body)
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "Page not found", 404)
			return
		}

		tmpl, err := template.ParseFiles("./layout.html", "./home.html")
		if err != nil {
			http.Error(w, "Error parsing template", 500)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(200)

		tmpl.ExecuteTemplate(w, "layout", "")

	})

	mux.HandleFunc("GET /paste/{id}/info", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		d, exist := ids[id]

		if !exist {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		resp := map[string]string{
			"created":        d.CreatedAt,
			"Content length": strconv.Itoa(len(d.Body)),
		}

		json.NewEncoder(w).Encode(resp)
	})

	mux.HandleFunc("GET /paste/{id}/view", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		d, exist := ids[id]

		if !exist {
			http.NotFound(w, r)
			return
		}

		tmpl, err := template.ParseFiles("./layout.html", "./paste_detail.html")
		if err != nil {
			http.Error(w, "Error parsing template", 500)
			fmt.Println(err)
			return
		}

		fmt.Println(d)
		tmpl.ExecuteTemplate(w, "layout", d)

	})

	http.ListenAndServe(
		":8080",
		mux,
	)

}
