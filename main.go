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
		return "", fmt.Errorf("name cant be empty")
	}

	s.count++
	return fmt.Sprintf("heello %s (request #%d)", name, s.count), nil
}

func greetHandler(g Greeter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("coolname")

		msg, err := g.Greet(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, msg)
	}
}

func main() {
	service := &GreetCounter{}

	http.HandleFunc("/greet", greetHandler(service))
	fmt.Println("server running on port 8080")
	http.ListenAndServe(":8080", nil)
}