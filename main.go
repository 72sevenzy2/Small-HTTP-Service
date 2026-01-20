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