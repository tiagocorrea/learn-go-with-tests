package main

import (
	"net/http"

	"github.com/tiagocorrea/go-specs-greet/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		panic(err)
	}
}
