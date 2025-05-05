package model

type RequestLeaveInstructor struct {
	BaseLeaveRequest
	InstructorCode string `gorm:"not null"`
}
