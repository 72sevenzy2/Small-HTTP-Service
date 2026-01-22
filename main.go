package main

import (
	"fmt"
	"net/http"
)

type Greeter interface {
	Greet(name string) (string, error)
}

type GreetCounter struct {
	count int
}

func (s *GreetCounter) Greet(name string) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name cannot be empty")
	}
	s.count++

	fmt.Println(s.count)

	return fmt.Sprintf("hello %s, greet number %d", name, s.count), nil
}

func GreetHandler(g Greeter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			name := r.URL.Query().Get("name")

			msg, err := g.Greet(name)

			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			fmt.Fprintln(w, msg)
		}
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "200")
}

func main() {
	service := &GreetCounter{}

	http.HandleFunc("/greet", GreetHandler(service))
	http.HandleFunc("/health", healthHandler)

	http.ListenAndServe(":8080", nil)
	fmt.Println("port listening on 8080")
}