package model

import (
	"time"
	"gorm.io/gorm"
)

type RequestLeaveInstructor struct {
	gorm.Model
	InstructorCode string    `gorm:"not null"` 
	Status    string    `gorm:"default:Pending"`
	LeaveType string  
	Reason    string  
	LeaveDate time.Time
}
