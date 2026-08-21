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

type MyHandler struct{}

func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, fibo(fiboCount))
	fiboCount++
}

func main() {
	handler := &MyHandler{}
	http.ListenAndServe(":8080", handler)
}
