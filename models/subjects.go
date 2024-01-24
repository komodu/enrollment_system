package models

import (
	"github.com/uadmin/uadmin"
)

type Subjects struct {
	uadmin.Model
	Name string `uadmin:"required"`
	//YearLevel YearLevel
}
