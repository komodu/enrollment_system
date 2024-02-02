package models

import (
	"github.com/uadmin/uadmin"
)

type CollegeCourses struct {
	uadmin.Model
	Course string `uadmin:"required"`
	//Subject string `uadmin:"required"`
	// YearLevel   YearLevel
	// YearLevelID uint
	//SchoolsID uint
	// CourseSHS CourseSHS `uadmin:"list_include"`
	//CourseSHSList string      `uadmin:"read_only"`
	// CourseSHSID uint
	// YearLevel     YearLevel
	// YearLevelID   uint
	//Subjects   Subjects `uadmin:"list_exclude"`
	//SubjectsID uint
	//StudentID  uint
}

func (s *CollegeCourses) String() string {
	return s.Course
}

// // Save !
// func (i *CollegeCourses) Save() {
// 	// Add a new string array type variable called categoryList
// 	CourseSHSList := []string{}

// 	// Append every element to the categoryList array
// 	for c := range i.CourseSHS {
// 		CourseSHSList = append(CourseSHSList, i.CourseSHS[c].Name)
// 	}

// 	// Concatenate the categoryList to a single string separated by comma
// 	joinList := strings.Join(CourseSHSList, ", ")

// 	// Store the joined string to the CategoryList field
// 	i.CourseSHSList = joinList

// 	// Save it to the database
// 	uadmin.Save(i)
// }
