package main

import (
	"log"
	"net/http"

	acceptancetests "github.com/tiagocorrea/go-graceful-shutdown/acceptance-tests"
)

func main() {
	server := &http.Server{Addr: ":8081", Handler: http.HandlerFunc(acceptancetests.SlowHandler)}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
