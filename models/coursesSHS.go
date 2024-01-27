package models

type CourseSHS int

func (CourseSHS) ABM() CourseSHS {
	return 1
}

func (CourseSHS) STEM() CourseSHS {
	return 2
}

func (CourseSHS) HUMSS() CourseSHS {
	return 3
}
func (CourseSHS) GAS() CourseSHS {
	return 4
}
func (CourseSHS) TVL() CourseSHS {
	return 5
}
