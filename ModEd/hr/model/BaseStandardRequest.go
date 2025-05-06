package model

import (
	"gorm.io/gorm"
)

// BaseStandardRequest holds fields common to Resignation and Raise requests
type BaseStandardRequest struct {
	gorm.Model
	Reason string `gorm:"type:text"`
	Status string `gorm:"default:Pending"`
}

// SetStatus implements RequestStatus.
func (b *BaseStandardRequest) SetStatus(status string) {
	b.Status = status
}

// SetReason implements RequestStatus.
func (b *BaseStandardRequest) SetReason(reason string) {
	b.Reason = reason
}

func (b *BaseStandardRequest) ApplyStatus(action Action, reason string) error {
	return ApplyStatus(b, action, reason)
}
