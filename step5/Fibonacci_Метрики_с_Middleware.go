package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

var fiboCount int
var mt sync.Mutex

func fibo(n int) int {
	if n == 1 {
		return 1
	} else if n <= 0 {
		return 0
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			mt.Lock()
			defer mt.Unlock()
			next.ServeHTTP(w, r)
			fiboCount++
			return
		}

		next.ServeHTTP(w, r)
		return
	})

}

func FiboHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.Itoa(fibo(fiboCount))))
}

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("rpc_duration_milliseconds_count %d", fiboCount)))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", FiboHandler)
	mux.HandleFunc("/metrics", MetricsHandler)

	handler := Metrics(mux)

	http.ListenAndServe(":8080", handler)
}
