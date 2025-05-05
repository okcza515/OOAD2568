package model

import "fmt"

type RequestLeaveStudent struct {
	BaseLeaveRequest
	StudentCode string `gorm:"not null"`
}

func (r *RequestLeaveStudent) ApplyStatus(action, reason string) error {
	switch action {
	case "approve":
		r.Status = action
	case "reject":
		r.Status = action
		r.Reason = reason
	default:
		return fmt.Errorf("invalid action: %q", action)
	}
	return nil
}
