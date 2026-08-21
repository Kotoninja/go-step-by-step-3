package main

import (
	"fmt"
	"net/http"
	"regexp"
)

func main() {
	fmt.Println(regexp.Match(`([A-Za-z])\w+`, []byte("Джон")))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name != "" {
			fmt.Fprintf(w, "hello stranger")
		} else if boo, _ := regexp.Match(`([A-Za-z])\w+`, []byte(name)); !boo {
			fmt.Fprintf(w, "hello dirty hacker")
		} else {
			fmt.Fprintf(w, ("hello " + name))
		}
	})

	http.ListenAndServe(":8080", nil)
}
