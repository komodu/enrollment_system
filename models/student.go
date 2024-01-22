package main

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Student struct {
	uadmin.Model
	FName        string
	LName        string
	Age          int
	PSA          bool
	GoodMoral    bool
	Parents      int
	Course       string
	DateEnrolled time.Time
}
