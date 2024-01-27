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
	CollegeCoursesList string           `uadmin:"list_exclude;read_only"`
	CollegeCoursesID   uint
	YearLevel          []YearLevel `uadmin:"list_exclude" gorm:"many2many:-"`
	YearLevelID        uint
	YearLevelList      string `uadmin:"read_only"`
	StudentID          uint   //`uadmin:"help:Students taking the subject;list_exclude"`
}

func (s *Subjects) String() string {
	return s.Name
}

// Save !
func (i *Subjects) Save() {
	// Add a new string array type variable called categoryList
	CollegeCoursesList := []string{}
	YearLevelList := []string{}
	// Append every element to the categoryList array
	for c := range i.CollegeCourses {
		CollegeCoursesList = append(CollegeCoursesList, i.CollegeCourses[c].Course)
	}
	for b := range i.YearLevel {
		YearLevelList = append(YearLevelList, i.YearLevel[b].Year)
	}
	// Concatenate the categoryList to a single string separated by comma
	joinList1 := strings.Join(CollegeCoursesList, ", ")
	joinList2 := strings.Join(YearLevelList, ", ")
	// Store the joined string to the CategoryList field
	i.CollegeCoursesList = joinList1
	i.YearLevelList = joinList2
	// Save it to the database
	uadmin.Save(i)

}
