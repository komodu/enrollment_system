package models

import (
	"github.com/uadmin/uadmin"
)

type CourseSHS struct {
	uadmin.Model
	Name string
	//CollegeCourses   CollegeCourses
	//CollegeCoursesID uint `uadmin:"list_exclude"`
}
