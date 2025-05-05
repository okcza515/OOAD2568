package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

// BaseStandardRequest holds fields common to Resignation and Raise requests
type BaseStandardRequest struct {
	gorm.Model
	Reason string `gorm:"type:text"`
	Status string `gorm:"default:Pending"`
}

func (b *BaseStandardRequest) ApplyStatus(action, reason string) error {
	switch action {
	case "approve":
		b.Status = action
	case "reject":
		b.Status = action
		b.Reason = reason
	default:
		return fmt.Errorf("invalid action: %q", action)
	}
	return nil
}

// BaseLeaveRequest holds fields common to Leave requests
type BaseLeaveRequest struct {
	gorm.Model
	Status    string `gorm:"default:Pending"`
	LeaveType string
	Reason    string `gorm:"type:text"`
	LeaveDate time.Time
}

func (b *BaseLeaveRequest) ApplyStatus(action, reason string) error {
	switch action {
	case "approve":
		b.Status = action
	case "reject":
		b.Status = action
		b.Reason = reason
	default:
		return fmt.Errorf("invalid action: %q", action)
	}
	return nil
}
