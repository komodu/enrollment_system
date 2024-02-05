package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// MainHandler is the main handler for the login system.
func MainHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called "/login_system/"
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/system/")
	page := strings.TrimSuffix(r.URL.Path, "/")
	// uadmin.Trail(uadmin.DEBUG, "\npage1: ", page)
	c := map[string]interface{}{}

	switch page {
	case "home": //Name of HTML
		c = HomePage(w, r)
	case "enroll": //Name of HTML
		c = StudentForm(w, r)
	case "test": //Name of HTML
		c = TestPage(w, r)
	case "checker": //Name of HTML
		c = StudentHandler(w, r)
	default:
		page = "home"
	}
	// uadmin.Trail(uadmin.DEBUG, "\npage2:", page)
	c["Page"] = page

	Rendering(w, r, page, c)

}

func Rendering(w http.ResponseWriter, r *http.Request, page string, context map[string]interface{}) {
	templateList := []string{}
	templateList = append(templateList, "./templates/base.html")
	//uadmin.Trail(uadmin.DEBUG, "test page: %v\n", page)

	path := "./templates/" + page + ".html"
	uadmin.Trail(uadmin.DEBUG, path)
	templateList = append(templateList, path)

	uadmin.RenderMultiHTML(w, r, templateList, context)
}
