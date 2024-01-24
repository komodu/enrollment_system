package models

import (
	"math/rand"
	"time"

	"github.com/uadmin/uadmin"
)

type Student struct {
	uadmin.Model
	AccountNum int `uadmin:"read_only"`
	//AccountNumID uint
	FullName  string `uadmin:"help:Follow the format LastName, FirstName, Middle Initial"`
	Age       int    `uadmin:"pattern:^[0-9]*$;required"`
	Address   string `uadmin:"list_exclude;required"`
	Email     string `uadmin:"email;required"`
	PSA       string `uadmin:"file;list_exclude"`
	GoodMoral string `uadmin:"file;list_exclude"`
	Form137   string `uadmin:"file;list_exclude"`
	Guardian  string
	// //AvailableCollegeCourses

	Schools   Schools `uadmin:"help:Where will you enroll?"`
	SchoolsID uint

	SHSCourse   CourseSHS `uadmin:"help:SHS Strand you previously enrolled?"`
	SHSCourseID uint
	Courses     CollegeCourses `uadmin:"help:What course will you enroll?"`
	CoursesID   uint

	YearLevel   YearLevel `uadmin:"help:What year level?"`
	YearLevelID uint

	// Subjects     Subjects
	// SubjectsID   uint
	DateEnrolled time.Time //`uadmin:"read_only"`
	Enrolled     bool      // `uadmin:"approval"`
}

func (s *Student) Compute() int {
	if s.Enrolled {
		s.AccountNum = rand.Intn(999999)
	}
	return s.AccountNum
}

// AccountNumber --> 8digit YYMMmmss --> 24010623
// func generateAcctNo(x *Student) string {
// 	min := 100000
// 	max := 999999
// 	x.AccountNumber = "24" + strconv.Itoa(rand.Intn(max-min)+min)
// 	return x.AccountNumber
// }

func (s *Student) String() string {
	return s.FullName
}

// func (x *Student) Save() string {
// 	min := 100000
// 	max := 999999
// 	x.AccountNumber = "24" + strconv.Itoa(rand.Intn(max-min)+min)
// 	return x.AccountNumber
// }

// Validate function !
func (t Student) Validate() (errMsg map[string]string) {
	// Initialize the error messages
	errMsg = map[string]string{}
	// Get any records from the database that macthes the name of
	// this record and make sure the record is not the record we are
	// editing right now
	student := Student{}
	if uadmin.Count(&student, "full_name = ? AND account_number <> ?", t.FullName, t.AccountNum) != 0 {
		errMsg["full_name"] = "The Student Name is already in the system"
	}
	return
}

// func (x *Student) Compute() string {
// 	// a := time.Now().Year()
// 	// b := time.Now().Minute()
// 	c := time.Now().UnixMilli()
// 	result := strconv.Itoa(c)
// 	result = x.AccountNum
// 	return result
// }

// func (completename *Student) Save() string {
// 	completename.FullName = completename.FirstName + " " + completename.LastName
// 	return completename.FullName
// }
