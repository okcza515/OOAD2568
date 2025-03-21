package model

type CourseStatus int

const (
	ACTIVE CourseStatus = iota
	INACTIVE
)

var CourseStatusLabel = map[CourseStatus]string{
	ACTIVE:   "active",
	INACTIVE: "inactive",
}

func (c CourseStatus) String() string {
	return CourseStatusLabel[c]
}

func (c CourseStatus) IsValid() bool {
	_, ok := CourseStatusLabel[c]
	return ok
}
