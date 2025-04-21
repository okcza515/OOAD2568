package model

type RequestRaiseBuilder struct {
	req *RequestRaise
}

func NewRequestRaiseBuilder() *RequestRaiseBuilder {
	return &RequestRaiseBuilder{
		req: &RequestRaise{
			Status: "Pending", // Default status
		},
	}
}

func (b *RequestRaiseBuilder) WithInstructorCode(code string) *RequestRaiseBuilder {
	b.req.InstructorCode = code
	return b
}

func (b *RequestRaiseBuilder) WithReason(reason string) *RequestRaiseBuilder {
	b.req.Reason = reason
	return b
}

func (b *RequestRaiseBuilder) WithStatus(status string) *RequestRaiseBuilder {
	b.req.Status = status
	return b
}

func (b *RequestRaiseBuilder) WithTargetSalary(targetSalary int) *RequestRaiseBuilder {
	b.req.TargetSalary = targetSalary
	return b
}

func (b *RequestRaiseBuilder) Build() *RequestRaise {
	return b.req
}