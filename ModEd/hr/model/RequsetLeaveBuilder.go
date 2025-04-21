package model
import (
	"time"	
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
			LeaveType: "ลาป่วย",
			Reason:    "เหตุผลการลา",
			LeaveDate: time.Now(),
		},
	}
}	
func NewRequestLeaveInstructorBuilder() *RequestLeaveInstructorBuilder{
	return &RequestLeaveInstructorBuilder{
		req: &RequestLeaveInstructor{
			Status:    "Pending",
			LeaveType: "ลาป่วย",
			Reason:    "เหตุผลการลา",
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

func (b *RequestLeaveStudentBuilder) WithLeaveDate(leaveDate time.Time) *RequestLeaveStudentBuilder {
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

func (b *RequestLeaveInstructorBuilder) WithLeaveDate(leaveDate time.Time) *RequestLeaveInstructorBuilder {
	b.req.LeaveDate = leaveDate
	return b
}
func (b *RequestLeaveInstructorBuilder) Build() *RequestLeaveInstructor {
	return b.req
}

 



