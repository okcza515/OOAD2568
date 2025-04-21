package utils

import (
	"ModEd/project/model"
	"errors"
)

type AdvisorBuilder struct {
	advisor model.Advisor
}

func NewAdvisorBuilder() *AdvisorBuilder {
	return &AdvisorBuilder{
		advisor: model.Advisor{},
	}
}

func (b *AdvisorBuilder) SetProjectID(projectId int) *AdvisorBuilder {
	b.advisor.SeniorProjectId = projectId
	return b
}

func (b *AdvisorBuilder) SetInstructorID(instructorId int) *AdvisorBuilder {
	b.advisor.InstructorId = instructorId
	return b
}

func (b *AdvisorBuilder) SetPrimaryStatus(isPrimary bool) *AdvisorBuilder {
	b.advisor.IsPrimary = isPrimary
	return b
}

func (b *AdvisorBuilder) Build() (*model.Advisor, error) {
	if b.advisor.SeniorProjectId == 0 {
		return nil, errors.New("project ID is required")
	}
	if b.advisor.InstructorId == 0 {
		return nil, errors.New("instructor ID is required")
	}
	return &b.advisor, nil
}
