package models

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/uadmin/uadmin"
)

type Student struct {
	uadmin.Model
	AccountNumber string `uadmin:"read_only"`
	// AccountNumberID uint

	FullName      string `uadmin:"read_only"`
	FirstName     string `uadmin:"list_exclude;required"`
	MiddleInitial string `uadmin:"list_exclude;required"`
	LastName      string `uadmin:"list_exclude;required"`
	Age           int    `uadmin:"pattern:^[0-9]*$;required"`
	Gender        Gender `uadmin:"required"`
	Address       string `uadmin:"list_exclude;required"`
	Email         string `uadmin:"email;required"`
	PSA           string `uadmin:"file;list_exclude"`
	GoodMoral     string `uadmin:"file;list_exclude"`
	Form137       string `uadmin:"file;list_exclude"`
	Guardian      string
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

// Gender Field !
type Gender int

func (Gender) Male() Gender {
	return 1
}

func (Gender) Female() Gender {
	return 2
}

func (t Student) Validate() (errMsg map[string]string) {
	// Initialize the error messages
	errMsg = map[string]string{}
	// Get any records from the database that macthes the name of
	// this record and make sure the record is not the record we are
	// editing right now
	student := Student{}
	if uadmin.Count(&student, "account_number <> ?", t.AccountNumber) == 0 {
		errMsg["account_number"] = "The Student is already in the system"
		uadmin.Trail(uadmin.DEBUG, errMsg)
		// uadmin.ALERT("")
	} else {
		uadmin.Trail(uadmin.DEBUG, "The Student is not in the system.")
	}
	return
}

func (s *Student) Save() {
	student := Student{}
	uadmin.Trail(uadmin.DEBUG, s.AccountNumber)

	if s.Enrolled {
		if uadmin.Count(&student, "account_number <> ?", s.AccountNumber) != 0 {
			s.AccountNumber = "24" + strconv.Itoa(rand.Intn(99)) + strconv.Itoa(rand.Intn(99)) + strconv.Itoa(rand.Intn(99)) + strconv.Itoa(rand.Intn(99))
			fmt.Println(s.AccountNumber)
		}
		//student := Student{}
		uadmin.Save(s)
	} else {
		s.AccountNumber = ""
		uadmin.Save(s)
	}
}
