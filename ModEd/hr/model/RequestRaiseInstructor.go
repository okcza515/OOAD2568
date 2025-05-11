package model

type RequestRaiseInstructor struct {
	BaseStandardRequest
	InstructorCode string `gorm:"not null"`
	TargetSalary   float64    `gorm:"not null"`
}
