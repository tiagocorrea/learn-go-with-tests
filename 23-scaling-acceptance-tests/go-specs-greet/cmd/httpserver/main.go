package main

import (
	"net/http"

	go_specs_greet "github.com/tiagocorrea/go-specs-greet"
)

func main() {
	handler := http.HandlerFunc(go_specs_greet.Handler)
	http.ListenAndServe(":8080", handler)
}
