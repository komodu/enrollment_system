package api

import (
	"net/http"
	"strings"
)

// Handler !
func Handler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called "/api/"
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api")
	r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	if strings.HasPrefix(r.URL.Path, "/test") {
		TodoListAPIHandler(w, r)
		return
	}
}
