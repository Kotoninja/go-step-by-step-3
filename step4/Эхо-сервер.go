package main

import (
	"fmt"
	"net/http"
)

func echoHandler(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	if msg == "" {
		msg = "empty"
	}
	fmt.Fprint(w, msg)
}

func main() {
	http.HandleFunc("/echo", echoHandler)

	http.ListenAndServe(":8080", nil)
}
