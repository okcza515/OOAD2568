package model

import (
	"time"
	"gorm.io/gorm"
)

type RequestLeaveStudent struct {
	gorm.Model
	StudentCode string    `gorm:"not null"` 
	Status    string    `gorm:"default:Pending"` 
	LeaveType string    
	Reason    string    
	LeaveDate time.Time 
}
