package model

type RequestResignationStudentBuilder struct {
	req *RequestResignationStudent
}

func NewRequestResignationStudentBuilder() *RequestResignationStudentBuilder {
	return &RequestResignationStudentBuilder{
		req: &RequestResignationStudent{
			Status: "Pending",
			// CreatedAt: time.Now(),
		},
	}
}

func (b *RequestResignationStudentBuilder) WithStudentID(id string) *RequestResignationStudentBuilder {
	b.req.StudentID = id
	return b
}

func (b *RequestResignationStudentBuilder) WithReason(reason string) *RequestResignationStudentBuilder {
	b.req.Reason = reason
	return b
}

func (b *RequestResignationStudentBuilder) WithStatus(status string) *RequestResignationStudentBuilder {
	b.req.Status = status
	return b
}

func (b *RequestResignationStudentBuilder) Build() *RequestResignationStudent {
	return b.req
}
