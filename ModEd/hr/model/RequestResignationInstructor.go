package model

type RequestResignationInstructor struct {
	BaseStandardRequest
	InstructorCode string `gorm:"not null"`
}
