package main

import (
	"github.com/uadmin/uadmin"
	"github.com/komodu/enrollment_system/models"

)

func main() {
	uadmin.Register{
		models.Student{},
	}
	uadmin.SiteName = "Enrollment System"
	//uadmin.BindIP = "127.0.0.1"
	uadmin.StartServer()
}
