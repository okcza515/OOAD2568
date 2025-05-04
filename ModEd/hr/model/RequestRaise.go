package model

type RequestRaise struct {
	BaseStandardRequest
	InstructorCode string `gorm:"not null"`
	TargetSalary   int    `gorm:"not null"`
}
