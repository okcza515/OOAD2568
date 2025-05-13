package model

import (
	"ModEd/core"
	"time"
)

// BaseLeaveRequest holds fields common to Leave requests
type BaseLeaveRequest struct {
	core.BaseModel
	Status    string `gorm:"default:Pending"`
	LeaveType string
	Reason    string `gorm:"type:text"`
	LeaveDate time.Time
}

// SetStatus implements RequestStatus.
func (b *BaseLeaveRequest) SetStatus(status string) {
	b.Status = status
}

// SetReason implements RequestStatus.
func (b *BaseLeaveRequest) SetReason(reason string) {
	b.Reason = reason
}

func (b *BaseLeaveRequest) ApplyStatus(action Action, reason string) error {
	return ApplyStatus(b, action, reason)
}
