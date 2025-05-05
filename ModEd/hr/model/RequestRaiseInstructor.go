package model

import "fmt"

type RequestRaiseInstructor struct {
    BaseStandardRequest
    InstructorCode string `gorm:"not null"`
    TargetSalary   int    `gorm:"not null"`
}

// ApplyStatus updates the status of the raise request based on the action and reason.
func (r *RequestRaiseInstructor) ApplyStatus(action, reason string) error {
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