package model

import "fmt"

type RequestResignationBuilder struct {
	req       interface{}
	isStudent bool
	err       error
}

// NewRequestResignationBuilder creates a new builder instance.
// Pass true for a student resignation request; false for an instructor.
func NewRequestResignationBuilder(isStudent bool) *RequestResignationBuilder {
	if isStudent {
		return &RequestResignationBuilder{
			req: &RequestResignationStudent{
				Status: "Pending",
			},
			isStudent: true,
		}
	}
	return &RequestResignationBuilder{
		req: &RequestResignationInstructor{
			Status: "Pending",
		},
		isStudent: false,
	}
}

// WithID sets the proper ID field based on resignation type.
func (b *RequestResignationBuilder) WithID(id string) *RequestResignationBuilder {
	if b.err != nil {
		return b
	}

	if id == "" {
		b.err = fmt.Errorf("id cannot be empty")
		return b
	}

	if b.isStudent {
		b.req.(*RequestResignationStudent).StudentCode = id
	} else {
		b.req.(*RequestResignationInstructor).InstructorCode = id
	}
	return b
}

// WithReason sets the resignation reason.
func (b *RequestResignationBuilder) WithReason(reason string) *RequestResignationBuilder {
	if b.err != nil {
		return b
	}

	if b.isStudent {
		b.req.(*RequestResignationStudent).Reason = reason
	} else {
		b.req.(*RequestResignationInstructor).Reason = reason
	}
	return b
}

// WithStatus sets the resignation status.
func (b *RequestResignationBuilder) WithStatus(status string) *RequestResignationBuilder {
	if b.err != nil {
		return b
	}

	if b.isStudent {
		b.req.(*RequestResignationStudent).Status = status
	} else {
		b.req.(*RequestResignationInstructor).Status = status
	}
	return b
}

// Build returns the built resignation request.
func (b *RequestResignationBuilder) Build() (interface{}, error) {
	return b.req, b.err
}
