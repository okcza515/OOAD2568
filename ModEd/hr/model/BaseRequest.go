package model

import (
	"time"

	"gorm.io/gorm"
)

// BaseStandardRequest holds fields common to Resignation and Raise requests
type BaseStandardRequest struct {
	gorm.Model
	Reason string `gorm:"type:text"`
	Status string `gorm:"default:Pending"`
}

// BaseLeaveRequest holds fields common to Leave requests
type BaseLeaveRequest struct {
	gorm.Model
	Status    string `gorm:"default:Pending"`
	LeaveType string
	Reason    string `gorm:"type:text"`
	LeaveDate time.Time
}
