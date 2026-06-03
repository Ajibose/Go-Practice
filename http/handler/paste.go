package handler

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"math/big"
	"net/http"
	"practice/http/model"
	"strconv"
	"time"
)

func CreatePaste(store *model.Storage) http.HandlerFunc {
	ids := store.Ids
	return func(w http.ResponseWriter, r *http.Request) {
		randomNum, _ := rand.Int(rand.Reader, big.NewInt(9000))
		id := randomNum.Int64() + 1000
		idString := strconv.Itoa(int(id))

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error reading body", err)
			http.Error(w, "failed to read body", http.StatusInternalServerError)
			return
		}

		d := model.Data{
			ID:        idString,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			Body:      string(body),
		}

		if len(body) == 0 {
			http.Error(w, "Body not provided", 400)
			return
		}

		ids[idString] = d
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(201)
		fmt.Fprintf(w, strconv.Itoa(int(id)))
	}
}

func GetAllPaste(store *model.Storage) http.HandlerFunc {
	ids := store.Ids
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("./templates/layout.html", "./templates/paste.html")
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error parsing templtate", 500)
			return
		}

		tmpl.ExecuteTemplate(w, "layout", ids)
	}
}

func GetPaste(store *model.Storage) http.HandlerFunc {
	ids := store.Ids
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		d, exist := ids[id]

		if !exist {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(200)
		fmt.Fprintf(w, d.Body)
	}
}

func GetPasteInfo(store *model.Storage) http.HandlerFunc {
	ids := store.Ids
	return func(w http.ResponseWriter, r *http.Request) {
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
	}
}

func ViewPaste(store *model.Storage) http.HandlerFunc {
	ids := store.Ids
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")

		d, exist := ids[id]

		if !exist {
			http.NotFound(w, r)
			return
		}

		tmpl, err := template.ParseFiles("./templates/layout.html", "./templates/paste_detail.html")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			fmt.Println(err)
			return
		}

		fmt.Println(d)
		tmpl.ExecuteTemplate(w, "layout", d)

	}
}
