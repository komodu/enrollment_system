package models

import "github.com/uadmin/uadmin"

// ==================================================+YEAR LEVEL ============================================
type YearLevel struct {
	uadmin.Model
	Year string `uadmin:"required"`

	SubjectsID uint `uadmin:"read_only"`
}

func (s *YearLevel) String() string {
	return s.Year
}
