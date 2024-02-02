package models

import "github.com/uadmin/uadmin"

// Category model ...
type Schools struct {
	uadmin.Model
	Name string `uadmin:"required"`
	Logo string `uadmin:"image"`
	//StudentID uint
}
