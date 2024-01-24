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
		models.Courses{},				
		models.Subjects{},
		models.YearLevel{},
		models.CourseSHS{},
	)
	uadmin.RegisterInlines(models.Schools{}, map[string]string{
		"Student":   "SchoolsID",
		"Courses":   "SchoolsID",
		"YearLevel": "SchoolsID",
	})
	uadmin.RegisterInlines(models.Subjects{}, map[string]string{
		"Courses":   "SubjectsID",
		"YearLevel": "SubjectsID",
	})

	uadmin.StartServer()
}
