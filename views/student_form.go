package views

import (
	"net/http"
	"strings"

	// Specify the username that you used inside github.com folder

	"github.com/komodu/enrollment_system/models"
	"github.com/uadmin/uadmin"
)

// StudentForm !
func StudentForm(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}
	//uadmin.Trail(uadmin.DEBUG, "enrollment_system")

	students := []models.Student{}
	school := []models.Schools{}
	gender := []models.Gender{}
	shsStrand := []models.CourseSHS{}

	uadmin.All(&students)
	for x := range students {
		uadmin.Preload(&students[x])
	}
	c["Title"] = "Enrollment System"

	// MAPPING FOR THE SCHOOL
	uadmin.All(&school)
	for x := range school {
		uadmin.Preload(&school[x])
	}
	//uadmin.Trail(uadmin.DEBUG, "Test school:")
	c["Schools"] = school

	// MAPPING FOR GENDER
	uadmin.All(&gender)
	for x := range gender {
		uadmin.Preload(&gender[x])
	}
	//uadmin.Trail(uadmin.DEBUG, "Test school:")
	c["Gender"] = gender

	// MAPPING FOR STRAND(SHS)
	uadmin.All(&shsStrand)
	for x := range shsStrand {
		uadmin.Preload(&shsStrand[x])
	}
	//uadmin.Trail(uadmin.DEBUG, "Test school:")
	c["SHS"] = shsStrand
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/enroll")
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/checker")
	return c
}
