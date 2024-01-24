package models

import (
	"time"

	//	"github.com/komodu/enrollment_system/models"
	"github.com/uadmin/uadmin"
)

type Student struct {
	uadmin.Model
	AccountNumber string
	FullName      string
	FirstName     string `uadmin:"list_exclude;required"`
	LastName      string `uadmin:"list_exclude;required"`
	Age           int    `uadmin:"pattern:^[0-9]*$;required"`
	Address       string `uadmin:"list_exclude;required"`
	Email         string `uadmin:"email;required"`
	PSA           bool
	GoodMoral     bool
	Guardian      string
	SHSCourse     CourseSHS
	SHSCourseID   uint
	//AvailableCollegeCourses

	Schools   Schools `uadmin:"help:Where will you enroll?"`
	SchoolsID uint

	Courses   Courses `uadmin:"help:What course will you enroll?"`
	CoursesID uint

	YearLevel   YearLevel `uadmin:"help:What year level?"`
	YearLevelID uint

	Subjects     Subjects
	SubjectsID   uint
	DateEnrolled time.Time `uadmin:"approval"` //`uadmin:"read_only"`
}

// AccountNumber --> 8digit YYMMmmss --> 24010623

func (s *Student) String() string {
	s.FullName = s.FirstName + " " + s.LastName
	return s.FullName
}

// func (x *Student) String() string {
// 	a := time.Now().Year()
// 	b := time.Now().Minute()
// 	c := time.Now().Second()
// 	x.AccountNumber := strconv.Itoa(a) + strconv.Itoa(b) + strconv.Itoa(c)

// 	return x.AccountNumber
// }

// func (completename *Student) Save() string {
// 	completename.FullName = completename.FirstName + " " + completename.LastName
// 	return completename.FullName
// }
