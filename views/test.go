package views

import (
	"net/http"
	"strings"
)

func TestPage(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}

	c["Title"] = "Test | Page"
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/test")
	//uadmin.RenderHTML(w, r, "./templates/test.html", context)
	return c
}
