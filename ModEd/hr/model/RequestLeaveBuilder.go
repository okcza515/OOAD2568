package model

import (
	"fmt"
	"time"
)

// Add Paid Leave

type RequestLeaveBuilder struct {
	req       interface{}
	err       error
	isStudent bool
}

// NewRequestLeaveBuilder creates a new builder.
// Pass true for a student leave; false for an instructor leave.
func NewRequestLeaveBuilder(isStudent bool) *RequestLeaveBuilder {
	if isStudent {
		return &RequestLeaveBuilder{
			req: &RequestLeaveStudent{
				Status:    "Pending",
				LeaveType: "sick",
				Reason:    "Reason for leave",
				LeaveDate: time.Now(),
			},
			isStudent: true,
		}
	}

	return &RequestLeaveBuilder{
		req: &RequestLeaveInstructor{
			Status:    "Pending",
			LeaveType: "sick",
			Reason:    "Reason for leave",
			LeaveDate: time.Now(),
		},
		isStudent: false,
	}
}

// WithID sets the proper ID field depending on the leave type.
func (b *RequestLeaveBuilder) WithID(id string) *RequestLeaveBuilder {
	if b.isStudent {
		b.req.(*RequestLeaveStudent).StudentCode = id
	} else {
		b.req.(*RequestLeaveInstructor).InstructorCode = id
	}
	return b
}

func (b *RequestLeaveBuilder) WithStatus(status string) *RequestLeaveBuilder {
	if b.isStudent {
		b.req.(*RequestLeaveStudent).Status = status
	} else {
		b.req.(*RequestLeaveInstructor).Status = status
	}
	return b
}

func (b *RequestLeaveBuilder) WithLeaveType(leaveType string) *RequestLeaveBuilder {
	if b.isStudent {
		b.req.(*RequestLeaveStudent).LeaveType = leaveType
	} else {
		b.req.(*RequestLeaveInstructor).LeaveType = leaveType
	}
	return b
}

func (b *RequestLeaveBuilder) WithReason(reason string) *RequestLeaveBuilder {
	if b.isStudent {
		b.req.(*RequestLeaveStudent).Reason = reason
	} else {
		b.req.(*RequestLeaveInstructor).Reason = reason
	}
	return b
}

func (b *RequestLeaveBuilder) WithLeaveDate(leaveDateStr string) *RequestLeaveBuilder {
	if b.err != nil {
		return b
	}
	t, err := time.Parse("2006-01-02", leaveDateStr)
	if err != nil {
		b.err = fmt.Errorf("invalid date format: %v", err)
		return b
	}
	if b.isStudent {
		b.req.(*RequestLeaveStudent).LeaveDate = t
	} else {
		b.req.(*RequestLeaveInstructor).LeaveDate = t
	}
	return b
}

// Build returns the constructed leave request and any error encountered.
// The caller will need to type-assert the result based on the original flag.
func (b *RequestLeaveBuilder) Build() (interface{}, error) {
	return b.req, b.err
}
