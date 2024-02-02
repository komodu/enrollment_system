package views

import (
	"net/http"
	"strings"

	// Specify the username that you used inside github.com folder

	"github.com/komodu/enrollment_system/models"
	"github.com/uadmin/uadmin"
)

// FriendsList !
func StudentForm(w http.ResponseWriter, r *http.Request) {
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
	//
	c := map[string]interface{}{}
	uadmin.Trail(uadmin.DEBUG, "enrollment_system")

	students := []models.Student{}
	school := []models.Schools{}
	gender := []models.Gender{}
	shsStrand := []models.CourseSHS{}
	uadmin.All(&students)

	for x := range students {
		uadmin.Preload(&students[x])
	}

	// c := map[string]interface{}{
	// 	"Title":    "Enrollment System",
	// 	"Students": students,
	// }
	c["Title"] = "Enrollment System1"

	// MAPPING FOR THE SCHOOL
	uadmin.All(&school)
	for x := range school {
		uadmin.Preload(&school[x])
	}
	uadmin.Trail(uadmin.DEBUG, "Test school:")
	c["Schools"] = school

	// MAPPING FOR GENDER
	uadmin.All(&gender)
	for x := range gender {
		uadmin.Preload(&gender[x])
	}
	uadmin.Trail(uadmin.DEBUG, "Test school:")
	c["Gender"] = gender

	// MAPPING FOR STRAND(SHS)
	uadmin.All(&shsStrand)
	for x := range shsStrand {
		uadmin.Preload(&shsStrand[x])
	}
	uadmin.Trail(uadmin.DEBUG, "Test school:")
	c["SHS"] = shsStrand

	uadmin.Trail(uadmin.DEBUG, "Test: ", c)
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/dashboard")
	uadmin.RenderHTML(w, r, "./templates/index.html", c)
	//return c
}
