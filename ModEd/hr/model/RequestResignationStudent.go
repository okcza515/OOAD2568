package model

import "fmt"

type RequestResignationStudent struct {
	BaseStandardRequest
	StudentCode string `gorm:"type:text;default:'';not null"`
}

func (r *RequestResignationStudent) ApplyStatus(action, reason string) error {
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
