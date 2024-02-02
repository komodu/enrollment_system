package main

import (
	// Specify the username that you used inside github.com folder
	"net/http"

	"github.com/komodu/enrollment_system/models"
	"github.com/komodu/enrollment_system/views"

	"github.com/uadmin/uadmin"
)

func main() {
	DBConfig()
	uadmin.Register(
		models.Student{},
		models.Schools{},
		models.Gender{},
		models.CourseSHS{},
		models.CollegeCourses{},
		models.Subjects{},
		models.YearLevel{},
	)

	http.HandleFunc("/enrollment_system/", uadmin.Handler(views.StudentForm))
	Server()
}

func DBConfig() {
	uadmin.Database = &uadmin.DBSettings{
		Type:     "mysql",
		Name:     "developers_task",
		User:     "root",
		Password: "Allen is Great 200%",
		Host:     "localhost",
		Port:     3306,
	}
}

func Server() {
	uadmin.RootURL = "/admin/"
	uadmin.Port = 1234
	uadmin.StartServer()
}
