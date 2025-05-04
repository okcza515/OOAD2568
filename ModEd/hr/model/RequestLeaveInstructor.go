package model

import "fmt"

type RequestLeaveInstructor struct {
	BaseLeaveRequest
	InstructorCode string `gorm:"not null"`
}

func (r *RequestLeaveInstructor) ApplyStatus(action, reason string) error {
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
