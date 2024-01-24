package models

import "github.com/uadmin/uadmin"

// ==================================================+YEAR LEVEL ============================================
type YearLevel struct {
	uadmin.Model
	Name       string
	SchoolsID  uint
	SubjectsID uint
}
