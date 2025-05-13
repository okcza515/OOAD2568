package model

import (
	"ModEd/core"
	"time"

	"github.com/go-playground/validator/v10"
)

// BaseLeaveRequest holds fields common to Leave requests
type BaseLeaveRequest struct {
	core.BaseModel
	Status    string `gorm:"default:Pending" validate:"required"`
	LeaveType string `validate:"required"`
	Reason    string `gorm:"type:text" validate:"-"`
	LeaveDate time.Time
}

func (baseLeaveRequest BaseLeaveRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(baseLeaveRequest); err != nil {
		return err
	}
	return nil
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
