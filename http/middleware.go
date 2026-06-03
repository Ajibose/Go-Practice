package main

import (
	"log"
	"net/http"
	"time"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("Request %v -> %v took %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}
