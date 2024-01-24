package models

import (
	"strings"

	"github.com/uadmin/uadmin"
)

type Subjects struct {
	uadmin.Model
	Name string `uadmin:"required"`
	//YearLevel YearLevel
	CollegeCourses     []CollegeCourses `uadmin:"list_exclude" gorm:"many2many:-"`
	CollegeCoursesList string           `uadmin:"read_only;list_exclude"`
	YearLevel          YearLevel        `uadmin:"list_exclude" gorm:"many2many:-"`
	YearLevelID        uint
}

// Save !
func (i *Subjects) Save() {
	// Add a new string array type variable called categoryList
	CollegeCoursesList := []string{}

	// Append every element to the categoryList array
	for c := range i.CollegeCourses {
		CollegeCoursesList = append(CollegeCoursesList, i.CollegeCourses[c].Course)
	}

	// Concatenate the categoryList to a single string separated by comma
	joinList := strings.Join(CollegeCoursesList, ", ")

	// Store the joined string to the CategoryList field
	i.CollegeCoursesList = joinList

	// Save it to the database
	uadmin.Save(i)
}

func (s *Subjects) String() string {
	return s.Name
}
