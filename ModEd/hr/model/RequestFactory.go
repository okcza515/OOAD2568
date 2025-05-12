package model

import (
	"fmt"
	"time"
)

// Role defines the type of user making a request.
type Role int

const (
	RoleStudent Role = iota
	RoleInstructor
)

type RequestType int

// RequestType defines the type of request.
const (
	RequestTypeLeave RequestType = iota
	RequestTypeResignation
	RequestTypeRaise
)

// CreateRequestParams holds the parameters for creating a request.
type CreateRequestParams struct {
	ID           string
	LeaveType    string // For Leave
	Reason       string
	DateStr      string
	TargetSalary float64    // For Raise
}

// requestCreator defines a function signature for creating a specific request.
type requestCreator func(params CreateRequestParams) (interface{}, error)

// --- Student Request Creators ---

func createStudentLeaveRequest(params CreateRequestParams) (interface{}, error) {
	if params.LeaveType == "" || params.Reason == "" || params.DateStr == "" {
		return nil, fmt.Errorf("missing parameters for student leave request (LeaveType, Reason, DateStr are required)")
	}
	t, err := time.Parse("2006-01-02", params.DateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format for student leave request ('%s'): %w", params.DateStr, err)
	}
	return &RequestLeaveStudent{
		BaseLeaveRequest: BaseLeaveRequest{
			Status:    "Pending",
			LeaveType: params.LeaveType,
			Reason:    params.Reason,
			LeaveDate: t,
		},
		StudentCode: params.ID,
	}, nil
}

func createStudentResignationRequest(params CreateRequestParams) (interface{}, error) {
	if params.Reason == "" {
		return nil, fmt.Errorf("missing Reason parameter for student resignation request")
	}
	return &RequestResignationStudent{
		BaseStandardRequest: BaseStandardRequest{
			Reason: params.Reason,
			Status: "Pending",
		},
		StudentCode: params.ID,
	}, nil
}

func createStudentRaiseRequest(params CreateRequestParams) (interface{}, error) {
	return nil, fmt.Errorf("students cannot request a raise")
}

// --- Instructor Request Creators ---

func createInstructorLeaveRequest(params CreateRequestParams) (interface{}, error) {
	if params.LeaveType == "" || params.Reason == "" || params.DateStr == "" {
		return nil, fmt.Errorf("missing parameters for instructor leave request (LeaveType, Reason, DateStr are required)")
	}
	t, err := time.Parse("2006-01-02", params.DateStr)
	if err != nil {
		return nil, fmt.Errorf("invalid date format for instructor leave request ('%s'): %w", params.DateStr, err)
	}
	return &RequestLeaveInstructor{
		BaseLeaveRequest: BaseLeaveRequest{
			Status:    "Pending",
			LeaveType: params.LeaveType,
			Reason:    params.Reason,
			LeaveDate: t,
		},
		InstructorCode: params.ID,
	}, nil
}

func createInstructorResignationRequest(params CreateRequestParams) (interface{}, error) {
	if params.Reason == "" {
		return nil, fmt.Errorf("missing Reason parameter for instructor resignation request")
	}
	return &RequestResignationInstructor{
		BaseStandardRequest: BaseStandardRequest{
			Reason: params.Reason,
			Status: "Pending",
		},
		InstructorCode: params.ID,
	}, nil
}

func createInstructorRaiseRequest(params CreateRequestParams) (interface{}, error) {
	if params.Reason == "" {
		return nil, fmt.Errorf("missing Reason parameter for instructor raise request")
	}
	return &RequestRaiseInstructor{
		BaseStandardRequest: BaseStandardRequest{
			Reason: params.Reason,
			Status: "Pending",
		},
		InstructorCode: params.ID,
		TargetSalary:   params.TargetSalary,
	}, nil
}

// requestCreatorsByRole holds all request creation functions, keyed by Role and then by RequestType.
var requestCreatorsByRole = map[Role]map[RequestType]requestCreator{
	RoleStudent: {
		RequestTypeLeave:       createStudentLeaveRequest,
		RequestTypeResignation: createStudentResignationRequest,
		RequestTypeRaise:       createStudentRaiseRequest,
	},
	RoleInstructor: {
		RequestTypeLeave:       createInstructorLeaveRequest,
		RequestTypeResignation: createInstructorResignationRequest,
		RequestTypeRaise:       createInstructorRaiseRequest,
	},
}

// RequestFactory is a unified factory for creating various types of requests.
type RequestFactory struct{}

// CreateRequest creates a specific request object based on the role, request type, and parameters.
func (f RequestFactory) CreateRequest(role Role, requestType RequestType, params CreateRequestParams) (interface{}, error) {
	if params.ID == "" {
		return nil, fmt.Errorf("ID parameter is required")
	}

	roleSpecificCreators, ok := requestCreatorsByRole[role]
	if !ok {
		return nil, fmt.Errorf("unknown role: %v", role)
	}

	creator, ok := roleSpecificCreators[requestType]
	if !ok {
		return nil, fmt.Errorf("unknown request type '%v' for role %v", requestType, role)
	}

	return creator(params)
}
