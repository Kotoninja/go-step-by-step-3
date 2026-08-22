package main

import (
	"encoding/json"
	// "fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	store *Store
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Store struct {
	pk    uint64
	users map[int]User
	mx    sync.Mutex
}

func NewStore() *Store {
	return &Store{pk: 1, mx: sync.Mutex{}, users: map[int]User{}}
}

func (s *Store) CreateUser(name string, age int) User {
	s.mx.Lock()
	defer func() {
		s.pk++
		s.mx.Unlock()
	}()

	newUser := User{ID: int(s.pk), Name: name, Age: age}
	s.users[int(s.pk)] = newUser

	return newUser
}

func (s *Store) GetUser(id int) (User, bool) {
	s.mx.Lock()
	defer s.mx.Unlock()

	user, ok := s.users[id]
	return user, ok
}

func (s *Store) GetAll() map[int]User {
	s.mx.Lock()
	defer s.mx.Unlock()
	return s.users
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	r.Body.Close()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	createdUser := store.CreateUser(newUser.Name, newUser.Age)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&createdUser)
	return
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")

	id, err := strconv.Atoi(idString)

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	user, ok := store.GetUser(id)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func loggingMiddleware(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rw := &responseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rw, r)
		end := time.Since(start)

		logger.Info("http request",
			slog.String("method", r.Method),
			slog.String("path", r.URL.Path),
			slog.Int("status", rw.statusCode),
			slog.Duration("duration", end))
	})
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	store = NewStore()

	slog.Info("Server start")
	mux := http.NewServeMux()

	mux.HandleFunc("/users", postHandler)
	mux.HandleFunc("/users/{id}", getHandler)

	handler := loggingMiddleware(logger, mux)

	err := http.ListenAndServe(":8080", handler)
	if err != nil {
		panic(err)
	}
}
