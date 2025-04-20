package model

type RequestResignationStudentBuilder struct {
	studentReq *RequestResignationStudent
}

type RequestResignationInstructorBuilder struct {
	instructorReq *RequestResignationInstructor
}

func NewRequestResignationStudentBuilder() *RequestResignationStudentBuilder {
	return &RequestResignationStudentBuilder{
		studentReq: &RequestResignationStudent{
			Status: "Pending",
		},
	}
}

func NewRequestResignationInstructorBuilder() *RequestResignationInstructorBuilder {
	return &RequestResignationInstructorBuilder{
		instructorReq: &RequestResignationInstructor{
			Status: "Pending",
		},
	}
}

// for student

func (b *RequestResignationStudentBuilder) WithStudentID(id string) *RequestResignationStudentBuilder {
	b.studentReq.StudentCode = id
	return b
}

func (b *RequestResignationStudentBuilder) WithReason(reason string) *RequestResignationStudentBuilder {
	b.studentReq.Reason = reason
	return b
}

func (b *RequestResignationStudentBuilder) WithStatus(status string) *RequestResignationStudentBuilder {
	b.studentReq.Status = status
	return b
}

func (b *RequestResignationStudentBuilder) Build() *RequestResignationStudent {
	return b.studentReq
}

// for instructor

func (b *RequestResignationInstructorBuilder) WithInstructorID(id string) *RequestResignationInstructorBuilder {
	b.instructorReq.InstructorCode = id
	return b
}

func (b *RequestResignationInstructorBuilder) WithReason(reason string) *RequestResignationInstructorBuilder {
	b.instructorReq.Reason = reason
	return b
}

func (b *RequestResignationInstructorBuilder) WithStatus(status string) *RequestResignationInstructorBuilder {
	b.instructorReq.Status = status
	return b
}

func (b *RequestResignationInstructorBuilder) Build() *RequestResignationInstructor {
	return b.instructorReq
}
