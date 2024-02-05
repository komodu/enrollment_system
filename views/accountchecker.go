package views

import (
	"net/http"

	// Specify the username that you used inside github.com folder

	"github.com/komodu/enrollment_system/models"
	"github.com/uadmin/uadmin"
)

// StudentForm !
func StudentHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {
	c := map[string]interface{}{}
	//uadmin.Trail(uadmin.DEBUG, "enrollment_system")
	students := []models.Student{}
	uadmin.All(&students)
	for i := range students {
		uadmin.Preload(&students[i])
		// c = append(c, map[string]interface{}{
		// 	"ID":             students[i].ID,
		// 	"Name":           students[i].FullName,
		// 	"CourseSHS":      students[i].CourseSHS,
		// 	"CollegeCourses": students[i].CollegeCourses,
		// 	"Schools":        students[i].Schools,
		// 	"YearLevel":      students[i].YearLevel,
		// 	"Gender":         students[i].Gender,
		// })
	}
	c["Students"] = students

	return c
}
