package main

import (
	"context"
	"net/http"
	"regexp"
)

type contextKey string

const nameKey contextKey = "name"

func Sanitize(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name != "" {
			matched, _ := regexp.MatchString("^[A-Za-z]+$", name)
			if !matched {
				w.Write([]byte("hello dirty hacker"))
				return
			}
		}
		next(w, r)
	})
}

func SetDefaultName(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "stranger"

		}
		ctx := context.WithValue(r.Context(), nameKey, name)
		next(w, r.WithContext(ctx))
	})
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name, ok := r.Context().Value(nameKey).(string)
	if !ok {
		name = "stranger"
	}

	w.Write([]byte("hello " + name))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", Sanitize(SetDefaultName(HelloHandler)))

	http.ListenAndServe(":8080", mux)
}
