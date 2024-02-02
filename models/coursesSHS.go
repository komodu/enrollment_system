package models

import (
	"strings"

	"github.com/uadmin/uadmin"
)

type CourseSHS struct {
	uadmin.Model
	Name              string
	CollegeCourses    []CollegeCourses `uadmin:"list_exclude" gorm:"many2many:-"`
	CollegeCourseList string           `uadmin:"read_only"`
	CollegeCoursesID  uint
} // Save !
func (i *CourseSHS) Save() {
	// Add a new string array type variable called categoryList
	CollegeCourseList := []string{}

	// Append every element to the categoryList array
	for c := range i.CollegeCourses {
		CollegeCourseList = append(CollegeCourseList, i.CollegeCourses[c].Course)
	}

	// Concatenate the categoryList to a single string separated by comma
	joinList := strings.Join(CollegeCourseList, ", ")

	// Store the joined string to the CategoryList field
	i.CollegeCourseList = joinList

	// Save it to the database
	uadmin.Save(i)
}
