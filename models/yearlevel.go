package models

import "github.com/uadmin/uadmin"

// ==================================================+YEAR LEVEL ============================================
type YearLevel struct {
	uadmin.Model
	Year string `uadmin:"required;search"`

	SubjectsID int `uadmin:"read_only;list_exclude"`
	StudentID  int
}

func (s *YearLevel) String() string {
	return s.Year
}
