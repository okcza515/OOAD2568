// MEP-1003 Student Recruitment
package controller

import (
	departmentCriteria "ModEd/recruit/controller/DepartmentCriteria"
	facultyCriteria "ModEd/recruit/controller/FacultyCriteria"
	roundCriteria "ModEd/recruit/controller/RoundCriteria"
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
		b.criteriaList = append(b.criteriaList, &roundCriteria.PortfolioCriteria{})
	case "Scholarship":
		b.criteriaList = append(b.criteriaList, &roundCriteria.ScholarshipCriteria{})
	case "Quota":
		b.criteriaList = append(b.criteriaList, &roundCriteria.QuotaCriteria{})
	case "Admission":
		b.criteriaList = append(b.criteriaList, &roundCriteria.AdmissionCriteria{})
	}
	return b
}

func (b *ApplicationCriteriaBuilder) AddFacultyCriteria(facultyName string) *ApplicationCriteriaBuilder {
	switch facultyName {
	case "Engineering":
		b.criteriaList = append(b.criteriaList, &facultyCriteria.EngineeringCriteria{})
	case "Science":
		b.criteriaList = append(b.criteriaList, &facultyCriteria.ScienceCriteria{})
	case "Architecture":
		b.criteriaList = append(b.criteriaList, &facultyCriteria.ArchitectureCriteria{})
	case "Business":
		b.criteriaList = append(b.criteriaList, &facultyCriteria.BusinessCriteria{})
	}
	return b
}

func (b *ApplicationCriteriaBuilder) AddDepartmentCriteria(departmentName string) *ApplicationCriteriaBuilder {
	switch departmentName {
	case "Computer Engineering":
		b.criteriaList = append(b.criteriaList, &departmentCriteria.ComputerEngineeringCriteria{})
	case "Mechanical Engineering":
		b.criteriaList = append(b.criteriaList, &departmentCriteria.MechanicalEngineeringCriteria{})
	case "Physics":
		b.criteriaList = append(b.criteriaList, &departmentCriteria.PhysicsCriteria{})
	case "Chemistry":
		b.criteriaList = append(b.criteriaList, &departmentCriteria.ChemistryCriteria{})
	case "Marketing":
		b.criteriaList = append(b.criteriaList, &departmentCriteria.MarketingCriteria{})
	case "Finance":
		b.criteriaList = append(b.criteriaList, &departmentCriteria.FinanceCriteria{})
	case "Interior Design":
		b.criteriaList = append(b.criteriaList, &departmentCriteria.InteriaCriteria{})
	case "Urban Planning":
		b.criteriaList = append(b.criteriaList, &departmentCriteria.UrbanCriteria{})
	}
	return b
}

func (b *ApplicationCriteriaBuilder) Build() Criteria {
	return &CompositeCriteria{criteriaList: b.criteriaList}
}
