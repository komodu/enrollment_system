package api

import (
	"net/http"

	// Specify the username that you used inside github.com folder
	"github.com/komodu/enrollment_system/models"
	"github.com/uadmin/uadmin"
)

// TodoListAPIHandler !
func TodoListAPIHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch all records in the database
	stud := []models.Student{}
	uadmin.All(&stud)

	// Accesses and fetches data from another model
	for t := range stud {
		uadmin.Preload(&stud[t])
	}

	// Return todo JSON object
	uadmin.ReturnJSON(w, r, stud)
}
