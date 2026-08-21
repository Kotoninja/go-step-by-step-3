package main

import (
	"fmt"
	"net/http"
)

var fiboCount int

func fibo(n int) int {
	if n == 1 {
		return 1
	} else if n <= 0 {
		return 0
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func main() {
	http.HandleFunc("/metric", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, fmt.Sprintf("rpc_duration_milliseconds_count %d", fiboCount))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprintln(w, fibo(fiboCount))
		fiboCount++
	})

	http.ListenAndServe(":8080", nil)
}
