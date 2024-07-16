package httpserver

import (
	"fmt"
	"net/http"

	"github.com/tiagocorrea/go-specs-greet/domain/interactions"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, interactions.Greet(name))
}
