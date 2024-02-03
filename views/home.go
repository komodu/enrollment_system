package views

import (
	"net/http"
	"strings"
)

func HomePage(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}

	c["Title"] = "Home | Page"
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/home")
	//uadmin.RenderHTML(w, r, "./templates/html/home.html", c)
	return c
}
