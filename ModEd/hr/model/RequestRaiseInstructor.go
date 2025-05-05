package model

type RequestRaiseInstructor struct {
	BaseStandardRequest
	InstructorCode string `gorm:"not null"`
	TargetSalary   int    `gorm:"not null"`
}
