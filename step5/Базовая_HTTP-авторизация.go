package main

import (
	"context"
	"fmt"
	"net/http"
)

func Authorization(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username == "" || password == "" {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintln(w, "Unauthorized")
			return
		}

		ctx := context.WithValue(r.Context(), "username", username)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func answerHandler(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)
	if !ok {
		username = "Guest"
	}

	fmt.Fprintf(w, "Welcome, %s!", username)

}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/answer/", Authorization(answerHandler))

	http.ListenAndServe(":8080", mux)
}
