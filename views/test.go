package views

import (
	"net/http"
	"strings"

	// Specify the username that you used inside github.com folder

	"github.com/komodu/enrollment_system/models"
	"github.com/uadmin/uadmin"
)

// FriendsList !
func TestForm(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	// r.URL.Path creates a new path called /friends
	// r.URL.Path = strings.TrimPrefix(r.URL.Path, "/friends")
	// r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")

	// Friends field inside the Context that will be used in Golang
	// HTML template
	// type Context struct {
	// 	Schools []map[string]interface{}
	// 	Student []map[string]interface{}
	// }

	// // Assigns Context struct to the c variable
	// c := Context{}

	// // Fetch Data from DB
	// school := []models.Schools{}
	// uadmin.All(&school)

	// for f := range school {
	// 	// Assigns the Friend records
	// 	c.Schools = append(c.Schools, map[string]interface{}{
	// 		"ID":   school[f].ID,
	// 		"Name": school[f].Name,
	// 		//   "Email": school[f].Email,
	// 	})
	// }

	// // Pass Friends data object to the specified HTML path

	uadmin.Trail(uadmin.DEBUG, "enrollment_system")

	students := []models.Student{}

	uadmin.All(&students)

	for s := range students {
		uadmin.Preload(&students[s])
	}

	c := map[string]interface{}{
		"Title":    "Enrollment System2",
		"Students": students,
	}
	uadmin.Trail(uadmin.DEBUG, "Test: ", c)
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/test")
	uadmin.RenderHTML(w, r, "./templates/test.html", c)
	return c
}
