package main

import "github.com/suborbital/sat/api/tinygo/runnable"

type Hello struct{}

func (h Hello) Run(input []byte) ([]byte, error) {
	return []byte("Hello, " + string(input)), nil
}

func main() {
	runnable.Use(Hello{})
}
