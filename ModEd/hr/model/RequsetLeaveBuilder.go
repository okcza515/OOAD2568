package model
import (
	"time"
	"fmt"
)
type RequestLeaveStudentBuilder struct {
	req *RequestLeaveStudent
}
type RequestLeaveInstructorBuilder struct {
	req *RequestLeaveInstructor
}

func NewRequestLeaveStudentBuilder() *RequestLeaveStudentBuilder {
	return &RequestLeaveStudentBuilder{
		req: &RequestLeaveStudent{
			Status:    "Pending",
			LeaveType: "sick",
			Reason:    "Reason for leave",
			LeaveDate: time.Now(),
		},
	}
}	
func NewRequestLeaveInstructorBuilder() *RequestLeaveInstructorBuilder{
	return &RequestLeaveInstructorBuilder{
		req: &RequestLeaveInstructor{
			Status:    "Pending",
			LeaveType: "sick",
			Reason:    "Reason for leave",
			LeaveDate: time.Now(),
		},
	}
}



func (b *RequestLeaveStudentBuilder) WithStudentID(id string) *RequestLeaveStudentBuilder {
	b.req.StudentCode = id
	return b
}

func (b *RequestLeaveStudentBuilder) WithStatus(status string) *RequestLeaveStudentBuilder {
	b.req.Status = status
	return b
}
func (b *RequestLeaveStudentBuilder) WithLeaveType(leaveType string) *RequestLeaveStudentBuilder {
	b.req.LeaveType = leaveType
	return b
}
func (b *RequestLeaveStudentBuilder) WithReason(reason string) *RequestLeaveStudentBuilder {
	b.req.Reason = reason
	return b
}

func (b *RequestLeaveStudentBuilder) WithLeaveDate(leaveDateStr string) *RequestLeaveStudentBuilder {
	leaveDate, err := time.Parse("2006-01-02", leaveDateStr)
	if err != nil {
		fmt.Printf("Invalid date format: %v\n", err)
		return b
	}
	b.req.LeaveDate = leaveDate
	return b
}

func (b *RequestLeaveStudentBuilder) Build() *RequestLeaveStudent {
	return b.req
}



func (b *RequestLeaveInstructorBuilder) WithInstructorID(id string) *RequestLeaveInstructorBuilder {
	b.req.InstructorCode = id
	return b
}

func (b *RequestLeaveInstructorBuilder) WithStatus(status string) *RequestLeaveInstructorBuilder {
	b.req.Status = status
	return b
}
func (b *RequestLeaveInstructorBuilder) WithLeaveType(leaveType string) *RequestLeaveInstructorBuilder {
	
	b.req.LeaveType = leaveType
	return b
}
func (b *RequestLeaveInstructorBuilder) WithReason(reason string) *RequestLeaveInstructorBuilder {
	b.req.Reason = reason
	return b
}

func (b *RequestLeaveInstructorBuilder) WithLeaveDate(leaveDateStr string) *RequestLeaveInstructorBuilder {
	leaveDate, err := time.Parse("2006-01-02", leaveDateStr)
	if err != nil {
		fmt.Printf("Invalid date format: %v\n", err)
		return b
	}
	b.req.LeaveDate = leaveDate
	return b
}
func (b *RequestLeaveInstructorBuilder) Build() *RequestLeaveInstructor {
	return b.req
}

 



