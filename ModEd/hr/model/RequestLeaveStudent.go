package model

type RequestLeaveStudent struct {
	BaseLeaveRequest
	StudentCode string `gorm:"not null"`
}
