package model

import (
	"fmt"
	"time"
)

type RequestLeaveProductInterface interface {
	GetID() string
	GetLeaveType() string
	GetReason() string
	GetLeaveDate() time.Time
	GetStatus() string
}

// RequestLeaveFactory is our simple factory for building leave requests.
type RequestLeaveFactory struct{}

// Create instantiates and returns a leave request based on the role.
func (f *RequestLeaveFactory) Create(role, id, leaveType, reason, dateStr string) (RequestLeaveProductInterface, error) {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format: %w", err)
	}
	switch role {
	case "student":
		return &RequestLeaveStudent{
			StudentCode: id,
			Status:      "Pending",
			LeaveType:   leaveType,
			Reason:      reason,
			LeaveDate:   t,
		}, nil
	case "instructor":
		return &RequestLeaveInstructor{
			InstructorCode: id,
			Status:         "Pending",
			LeaveType:      leaveType,
			Reason:         reason,
			LeaveDate:      t,
		}, nil
	default:
		return nil, fmt.Errorf("invalid role: %s", role)
	}
}
