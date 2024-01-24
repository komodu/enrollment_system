package models

import (
	"github.com/uadmin/uadmin"
)

type Courses struct {
	uadmin.Model
	Course  string `uadmin:"required"`
	Subject string `uadmin:"required"`
	//YearLevel YearLevel
	SchoolsID  uint
	SubjectsID uint
}
