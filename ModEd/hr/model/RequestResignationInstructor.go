package model

import "fmt"

type RequestResignationInstructor struct {
	BaseStandardRequest
	InstructorCode string `gorm:"not null"`
}

func (r *RequestResignationInstructor) ApplyStatus(action, reason string) error {
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
