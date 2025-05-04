package model

import (
	"fmt"
	"time"
)

// AbstractFactory interface
type AbstractFactory interface {
	CreateLeave(id, leaveType, reason, dateStr string) (interface{}, error)
	CreateResignation(id, reason string) (interface{}, error)
	CreateRaise(id, reason string, targetSalary int) (interface{}, error)
}

// Concrete factories implementing the AbstractFactory interface
type StudentFactory struct{}

// CreateLeave creates a leave request for a student
func (StudentFactory) CreateLeave(id, leaveType, reason, dateStr string) (interface{}, error) {
	t, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date: %w", err)
	}
	// Initialize embedded BaseLeaveRequest and specific fields
	return &RequestLeaveStudent{
		BaseLeaveRequest: BaseLeaveRequest{
			Status:    "Pending",
			LeaveType: leaveType,
			Reason:    reason,
			LeaveDate: t,
		},
		StudentCode: id,
	}, nil
}

func (StudentFactory) CreateResignation(id, reason string) (interface{}, error) {
	// Initialize embedded BaseStandardRequest and specific fields
	return &RequestResignationStudent{
		BaseStandardRequest: BaseStandardRequest{
			Reason: reason,
			Status: "Pending",
		},
		StudentCode: id,
	}, nil
}

func (StudentFactory) CreateRaise(id, reason string, targetSalary int) (interface{}, error) {
	return nil, fmt.Errorf("students cannot request a raise")
}

// Concrete factories implementing the AbstractFactory interface
type InstructorFactory struct{}

// CreateLeave creates a leave request for an instructor
func (InstructorFactory) CreateLeave(id, leaveType, reason, dateStr string) (interface{}, error) {
	t, err := time.Parse("02-01-2006", dateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date: %w", err)
	}
	// Initialize embedded BaseLeaveRequest and specific fields
	return &RequestLeaveInstructor{
		BaseLeaveRequest: BaseLeaveRequest{
			Status:    "Pending",
			LeaveType: leaveType,
			Reason:    reason,
			LeaveDate: t,
		},
		InstructorCode: id,
	}, nil
}

func (InstructorFactory) CreateResignation(id, reason string) (interface{}, error) {
	// Initialize embedded BaseStandardRequest and specific fields
	return &RequestResignationInstructor{
		BaseStandardRequest: BaseStandardRequest{
			Reason: reason,
			Status: "Pending",
		},
		InstructorCode: id,
	}, nil
}

func (InstructorFactory) CreateRaise(id, reason string, targetSalary int) (interface{}, error) {
	// Initialize embedded BaseStandardRequest and specific fields
	return &RequestRaise{
		BaseStandardRequest: BaseStandardRequest{
			Reason: reason,
			Status: "Pending",
		},
		InstructorCode: id,
		TargetSalary:   targetSalary,
	}, nil
}

// GetFactory returns the appropriate factory based on the role provided
func GetFactory(role string) (AbstractFactory, error) {
	switch role {
	case "student":
		return StudentFactory{}, nil
	case "instructor":
		return InstructorFactory{}, nil
	default:
		return nil, fmt.Errorf("unknown role %q", role)
	}
}
