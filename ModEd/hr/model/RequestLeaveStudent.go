package model

import (
	"gorm.io/gorm"
	"time"
)

type RequestLeaveStudent struct {
	gorm.Model
	StudentCode string `gorm:"not null"`
	Status      string `gorm:"default:Pending"`
	LeaveType   string
	Reason      string
	LeaveDate   time.Time
}

func (r *RequestLeaveStudent) GetID() string {
	return r.StudentCode
}
func (r *RequestLeaveStudent) GetLeaveType() string {
	return r.LeaveType
}
func (r *RequestLeaveStudent) GetReason() string {
	return r.Reason
}
func (r *RequestLeaveStudent) GetLeaveDate() time.Time {
	return r.LeaveDate
}
func (r *RequestLeaveStudent) GetStatus() string {
	return r.Status
}
