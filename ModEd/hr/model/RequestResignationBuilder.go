package model
import (
	"time"	
)
type RequestResignationBuilder struct {
	req *RequestResignation
}

func NewRequestResignationBuilder() *RequestResignationBuilder {
	return &RequestResignationBuilder{
		req: &RequestResignation{
			Status:    "Pending",
			CreatedAt: time.Now(),
		},
	}
}

func (b *RequestResignationBuilder) WithStudentID(id string) *RequestResignationBuilder {
	b.req.StudentID = id
	return b
}

func (b *RequestResignationBuilder) WithReason(reason string) *RequestResignationBuilder {
	b.req.Reason = reason
	return b
}

func (b *RequestResignationBuilder) WithStatus(status string) *RequestResignationBuilder {
	b.req.Status = status
	return b
}

func (b *RequestResignationBuilder) Build() *RequestResignation {
	return b.req
}