package model

import "fmt"

type RequestResignationProductInterface interface {
	GetID() string
	GetReason() string
	GetStatus() string
}

type RequestResignationFactory struct{}

func (f *RequestResignationFactory) Create(role, id, reason string) (RequestResignationProductInterface, error) {
	switch role {
	case "student":
		return &RequestResignationStudent{
			StudentCode: id,
			Reason:      reason,
			Status:      "Pending",
		}, nil
	case "instructor":
		return &RequestResignationInstructor{
			InstructorCode: id,
			Reason:         reason,
			Status:         "Pending",
		}, nil
	default:
		return nil, fmt.Errorf("invalid role: %s", role)
	}
}
