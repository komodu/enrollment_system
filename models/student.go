package models

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/uadmin/uadmin"
)

type Student struct {
	uadmin.Model
	AccountNumber string `uadmin:"read_only"`

	// AccountNumberID uint
	AutoGenerateAccNum bool `uadmin:"read_only;help:This is Automatically turned On.;list_exclude"`

	FullName  string `uadmin:"read_only;list_exclude"`
	FirstName string `uadmin:"list_exclude"`
	//MiddleInitial string `uadmin:"list_exclude"`
	LastName  string `uadmin:"list_exclude"`
	Age       int    `uadmin:"pattern:^[0-9]*$;list_exclude"`
	Gender    Gender //`uadmin:"list_exclude"`
	Address   string `uadmin:"list_exclude"`
	Email     string `uadmin:"email;list_exclude"`
	PSA       string `uadmin:"file;list_exclude"`
	GoodMoral string `uadmin:"file;list_exclude"`
	Form137   string `uadmin:"file;list_exclude"`
	Guardian  string
	// //AvailableCollegeCourses

	Schools   Schools `uadmin:"help:Where will you enroll?"`
	SchoolsID uint

	CourseSHS   CourseSHS `uadmin:"help:SHS Strand you previously enrolled?"`
	CourseSHSID uint

	CollegeCourses   CollegeCourses `uadmin:"help:What course will you enroll?"`
	CollegeCoursesID uint

	// //CollegeCoursesID uint `uadmin:"list_exclude"`

	YearLevel   YearLevel `uadmin:"help:What year level?"`
	YearLevelID uint

	DateEnrolled time.Time
	Enrolled     bool `uadmin:"help:"`
}

func (s *Student) String() string {
	return s.FullName
}

func (r *Student) Generate() string {
	r.AutoGenerateAccNum = true
	if r.AutoGenerateAccNum {
		ind1 := r.DateEnrolled.Year()
		ind2 := strconv.Itoa(rand.Intn(9)) + strconv.Itoa(rand.Intn(9)) + strconv.Itoa(rand.Intn(9)) + strconv.Itoa(rand.Intn(9)) + strconv.Itoa(rand.Intn(9))

		r.AccountNumber = strconv.Itoa(ind1) + ind2
		return r.AccountNumber
	}
	uadmin.Save(r)
	uadmin.Trail(uadmin.DEBUG, r.AccountNumber)
	return r.AccountNumber
}

// Validate function !
func (t Student) Validate() (errMsg map[string]string) {
	// Initialize the error messages
	errMsg = map[string]string{}
	// Get any records from the database that maches the name of
	// this record and make sure the record is not the record we are
	// editing right now
	// stud := Student{}
	// if uadmin.Count(&stud, "name = ? AND id <> ?", t.CourseSHS, t.CollegeCourses.CourseSHS) != 0 {
	// 	errMsg["Name"] = "This todo name is already in the system"
	// }
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
	//s.AutoGenerateAccNum = true

	if s.AccountNumber == "" {
		uadmin.Trail(uadmin.DEBUG, s.AccountNumber)
		s.FullName = s.LastName + " " + s.FirstName + " "
		//errMsg["account_number"] = "The Student is already in the system"
		//t.AccountNumber = ind + strconv.Itoa(x) + strconv.Itoa(y)
		for {
			s.AccountNumber = s.Generate()
			if uadmin.Count(&s, "account_number = ?", s.AccountNumber) == 0 {
				uadmin.Save(s)
				uadmin.Trail(uadmin.DEBUG, "The Account number is not in the server. Number added.")
				uadmin.Trail(uadmin.DEBUG, s.AccountNumber)
				break
				// uadmin.ALERT("")
			} else {
				//uadmin.Save(s)

				uadmin.Trail(uadmin.DEBUG, "The Account number has duplicate. Number not added. Regenerating new number ")
				//s.AccountNumber = s.Generate()
				uadmin.Trail(uadmin.DEBUG, s.AccountNumber)
			}
		}
		// uadmin.ALERT("")
	} else {
		uadmin.Save(s)
	}
}
