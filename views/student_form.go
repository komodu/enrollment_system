package views

import (
	"net/http"

	// Specify the username that you used inside github.com folder

	"github.com/komodu/enrollment_system/models"
	"github.com/uadmin/uadmin"
)

// FriendsList !
func StudentForm(w http.ResponseWriter, r *http.Request) map[string]interface{} {
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
	return c
}
