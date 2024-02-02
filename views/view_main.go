package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// MainHandler is the main handler for the login system.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called "/login_system/"
	// r.URL.Path = strings.TrimPrefix(r.URL.Path, "//")
	page := strings.TrimSuffix(r.URL.Path, "/")
	uadmin.Trail(uadmin.DEBUG, "\npage1: ", page)
	c := map[string]interface{}{}

	switch page {
	case "index": //Name of HTML
		c = StudentForm(w, r)
	case "test": //Name of HTML
		c = TestForm(w, r)
	default:
		page = "index"
		c = StudentForm(w, r)
	}
	uadmin.Trail(uadmin.DEBUG, "\npage2:", page)
	c["Page"] = page

	Rendering(w, r, page, c)

}

func Rendering(w http.ResponseWriter, r *http.Request, page string, context map[string]interface{}) {
	templateList := []string{}
	templateList = append(templateList, "./templates/index.html")

	path := "./templates/" + page + ".html"
	templateList = append(templateList, path)

	uadmin.RenderMultiHTML(w, r, templateList, context)
}
