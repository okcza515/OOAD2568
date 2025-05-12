// MEP-1003 Student Recruitment
package controller

import (
	departmentCriteria "ModEd/recruit/controller/DepartmentCriteria"
	facultyCriteria "ModEd/recruit/controller/FacultyCriteria"
)

type ApplicationCriteriaBuilder struct {
	criteriaList []Criteria
}

func NewApplicationCriteriaBuilder() *ApplicationCriteriaBuilder {
	return &ApplicationCriteriaBuilder{}
}

func (b *ApplicationCriteriaBuilder) AddRoundCriteria(roundName string) *ApplicationCriteriaBuilder {
	switch roundName {
	case "Portfolio":
		b.criteriaList = append(b.criteriaList, &PortfolioCriteria{})
	case "Scholarship":
		b.criteriaList = append(b.criteriaList, &ScholarshipCriteria{})
	case "Quota":
		b.criteriaList = append(b.criteriaList, &QuotaCriteria{})
	case "Admission":
		b.criteriaList = append(b.criteriaList, &AdmissionCriteria{})
	}
	return b
}

func (b *ApplicationCriteriaBuilder) AddFacultyCriteria(facultyName string) *ApplicationCriteriaBuilder {
	switch facultyName {
	case "Engineering":
		b.criteriaList = append(b.criteriaList, &facultyCriteria.EngineeringCriteria{})
	}
	return b
}

func (b *ApplicationCriteriaBuilder) AddDepartmentCriteria(departmentName string) *ApplicationCriteriaBuilder {
	switch departmentName {
	case "Computer Engineering":
		b.criteriaList = append(b.criteriaList, &departmentCriteria.ComputerEngineeringCriteria{})
	case "Mechanical Engineering":
		b.criteriaList = append(b.criteriaList, &departmentCriteria.MechanicalEngineeringCriteria{})
	}

	return b
}

func (b *ApplicationCriteriaBuilder) Build() Criteria {
	return &CompositeCriteria{criteriaList: b.criteriaList}
}
