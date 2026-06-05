package main

import (
	"bytes"
	"crypto/rand"
	_ "embed"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"practice/ascii-art-web-practice/ascii"
	"sync"
	"time"
)

//go:embed banner/standard.txt
var standard string

//go:embed banner/shadow.txt
var shadow string

//go:embed banner/thinkertoy.txt
var thinkertoy string

var m sync.Map

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", home)
	mux.HandleFunc("POST /ascii-art", asciiArt)

	log.Println("Server listening on http://localhost:8080")
	err := http.ListenAndServe(":8080", logger(mux))
	if err != nil {
		log.Fatal("Unable to start server on port 8080", err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("./templates/index.html")
	if errors.Is(err, os.ErrNotExist) {
		http.Error(w, "Template file not exist", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Println("Error parsing template", err)
		http.Error(w, "Error Parsing Template", http.StatusInternalServerError)
		return
	}

	var data any

	c, err := r.Cookie("id")
	if err == nil {
		data, _ = m.LoadAndDelete(c.Value)
		c.MaxAge = -1
		http.SetCookie(w, c)
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		http.Error(w, "Error Executing Template", http.StatusInternalServerError)
		return
	}
	buf.WriteTo(w)
}

func asciiArt(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Form Data", http.StatusBadRequest)
		return
	}

	banner := r.FormValue("banner")
	text := r.FormValue("text")

	if banner == "" || text == "" {
		http.Error(w, "Empty Field Provided", http.StatusBadRequest)
		return
	}

	var style string

	switch banner {
	case "standard":
		style = standard
	case "shadow":
		style = shadow
	case "thinkertoy":
		style = thinkertoy
	default:
		http.Error(w, "Specified Banner does not exist", http.StatusNotFound)
		return
	}

	asciiArt := ascii.GenerateAsciiArt(style, text)

	id := make([]byte, 16)
	_, err = rand.Read(id)
	if err != nil {
		http.Error(w, "Error generating key", http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{
		Name:     "id",
		Value:    fmt.Sprintf("%x", id),
		Path:     "/",
		Expires:  time.Now().Add(time.Hour),
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteStrictMode,
	}

	m.Store(fmt.Sprintf("%x", id), asciiArt)

	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Request %v %v -> took %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}
