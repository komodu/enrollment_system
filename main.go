package main

import (
	// Specify the username that you used inside github.com folder
	"github.com/komodu/enrollment_system/models"

	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.Register(
		models.Student{},
		models.Schools{},
		models.CourseSHS{},
		models.CollegeCourses{},
		models.Subjects{},
		models.YearLevel{},
	)

	uadmin.RegisterInlines(
		models.Schools{},
		map[string]string{
			"Student": "SchoolsID",
		},
	)
	uadmin.RegisterInlines(
		models.YearLevel{},
		map[string]string{
			"Student": "YearLevelID",
		},
	)
	uadmin.RegisterInlines(
		models.CollegeCourses{},
		map[string]string{
			"CourseSHS": "CollegeCoursesID",
		},
	)

	uadmin.StartServer()
}
