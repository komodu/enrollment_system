package models

import (
	"strconv"
	"time"

	"github.com/uadmin/uadmin"
)

type Student struct {
	uadmin.Model
	AccountNumber string `uadmin:"read_only"`
	// AccountNumberID uint

	FullName  string `uadmin:"read_only"`
	FirstName string `uadmin:"list_exclude"`
	//MiddleInitial string `uadmin:"list_exclude"`
	LastName  string `uadmin:"list_exclude"`
	Age       int    `uadmin:"pattern:^[0-9]*$"`
	Gender    Gender `uadmin:"required"`
	Address   string `uadmin:"list_exclude"`
	Email     string `uadmin:"email"`
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

func (s *Student) String() string {
	return s.FullName
}

func (t *Student) Validate() (errMsg map[string]string) {
	// Initialize the error messages
	errMsg = map[string]string{}
	ind := "2400"
	x := 1000
	y := uadmin.Count(&t, "account_number <> ?", t.AccountNumber)
	// Get any records from the database that macthes the name of
	// this record and make sure the record is not the record we are
	// editing right now
	if t.Enrolled {

		if uadmin.Count(&t, "account_number <> ?", t.AccountNumber) != 0 {
			errMsg["account_number"] = "The Student is already in the system"
			t.AccountNumber = ind + strconv.Itoa(x) + strconv.Itoa(y)
			uadmin.Trail(uadmin.DEBUG, t.AccountNumber)
			// uadmin.ALERT("")
		} else {
			//t.AccountNumber = strconv.Itoa(rand.Intn(10))
			t.AccountNumber = ind + strconv.Itoa(x)
			uadmin.Trail(uadmin.DEBUG, "The Student is not in the system.")
		}
		uadmin.Save(t)
	}
	return
}

// Gender Field !
type Gender int

func (Gender) Male() Gender {
	return 1
}

func (Gender) Female() Gender {
	return 2
}

func (s *Student) Save() {
	uadmin.Trail(uadmin.DEBUG, s.AccountNumber)
	s.FullName = s.LastName + " " + s.FirstName + " "
	uadmin.Save(s)
	// if s.Enrolled {
	// 	if uadmin.Count(&s, "account_number <> ?", s.AccountNumber) == 0 {
	// 		//fmt.Println(s.AccountNumber)
	// 		uadmin.Trail(uadmin.DEBUG, s.AccountNumber)
	// 		uadmin.Trail(uadmin.DEBUG, "No Duplicated Acc number found. Student successfully added.")
	// 		uadmin.Save(s)
	// 	}
	// 	//student := Student{}
	// } else {
	// 	s.AccountNumber = ""
	// 	uadmin.Trail(uadmin.DEBUG, "Duplicate found. Adding")
	// 	uadmin.Save(s)
	// }
}
