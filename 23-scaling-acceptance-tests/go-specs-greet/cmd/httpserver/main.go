package main

import (
	"log"
	"net/http"

	"github.com/tiagocorrea/go-specs-greet/adapters/httpserver"
)

func main() {
	handler := httpserver.NewHandler()
	log.Fatal(http.ListenAndServe(":8080", handler))
}
