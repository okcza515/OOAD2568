package model

import (
	"gorm.io/gorm"
	"time"
)

type RequestLeaveInstructor struct {
	gorm.Model
	InstructorCode string `gorm:"not null"`
	Status         string `gorm:"default:Pending"`
	LeaveType      string
	Reason         string
	LeaveDate      time.Time
}

func (r RequestLeaveInstructor) GetID() string {
	return r.InstructorCode
}

func (r RequestLeaveInstructor) GetLeaveType() string {
	return r.LeaveType
}

func (r RequestLeaveInstructor) GetReason() string {
	return r.Reason
}

func (r RequestLeaveInstructor) GetLeaveDate() time.Time {
	return r.LeaveDate
}

func (r RequestLeaveInstructor) GetStatus() string {
	return r.Status
}
