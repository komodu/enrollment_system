package models

// ==================================================+YEAR LEVEL ============================================
type YearLevel int

func (YearLevel) FirstYear() YearLevel {
	return 1
}

func (YearLevel) SecondYear() YearLevel {
	return 2
}

func (YearLevel) ThirdYear() YearLevel {
	return 3
}
func (YearLevel) FourthYear() YearLevel {
	return 4
}
func (YearLevel) FifthYear() YearLevel {
	return 5
}
