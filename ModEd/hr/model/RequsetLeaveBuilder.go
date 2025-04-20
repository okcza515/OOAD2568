package model
import (
	"time"	
)
type RequestLeaveBuilder struct {
	req *RequestLeave
}

func NewRequestLeaveBuilder() *RequestLeaveBuilder {
	return &RequestLeaveBuilder{
		req: &RequestLeave{
			CreatedAt: time.Now(),
		},
	}
}
func (b *RequestLeaveBuilder) WithStudentID(id string) *RequestLeaveBuilder {
	b.req.StudentID = id
	return b
}

func (b *RequestLeaveBuilder) WithStatus(status string) *RequestLeaveBuilder {
	b.req.Status = status
	return b
}
func (b *RequestLeaveBuilder) WithLeaveType(leaveType string) *RequestLeaveBuilder {
	b.req.LeaveType = leaveType
	return b
}
func (b *RequestLeaveBuilder) WithReason(reason string) *RequestLeaveBuilder {
	b.req.Reason = reason
	return b
}
func (b *RequestLeaveBuilder) WithRemarks(remarks string) *RequestLeaveBuilder {
	b.req.Remarks = remarks
	return b
}
func (b *RequestLeaveBuilder) WithApprovedBy(approvedBy string) *RequestLeaveBuilder {
	b.req.ApprovedBy = approvedBy
	return b
}
func (b *RequestLeaveBuilder) WithLeaveDate(leaveDate time.Time) *RequestLeaveBuilder {
	b.req.LeaveDate = leaveDate
	return b
}
func (b *RequestLeaveBuilder) WithApprovedAt(approvedAt time.Time) *RequestLeaveBuilder {
	b.req.ApprovedAt = approvedAt
	return b
}
func (b *RequestLeaveBuilder) Build() *RequestLeave {
	return b.req
}
func (b *RequestLeaveBuilder) WithCreatedAt(createdAt time.Time) *RequestLeaveBuilder {
	b.req.CreatedAt = createdAt
	return b
}
func (b *RequestLeaveBuilder) WithID(id uint) *RequestLeaveBuilder {
	b.req.ID = id
	return b
}


