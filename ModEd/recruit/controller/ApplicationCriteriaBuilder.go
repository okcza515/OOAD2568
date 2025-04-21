// MEP-1003 Student Recruitment
package controller

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
	}
	return b
}

func (b *ApplicationCriteriaBuilder) AddFacultyCriteria(facultyName string) *ApplicationCriteriaBuilder {
	switch facultyName {
	case "Engineering":
		b.criteriaList = append(b.criteriaList, &EngineeringCriteria{})
	}
	return b
}

func (b *ApplicationCriteriaBuilder) AddDepartmentCriteria(departmentName string) *ApplicationCriteriaBuilder {
	switch departmentName {
	case "Computer Engineering":
		b.criteriaList = append(b.criteriaList, &ComputerEngineeringCriteria{})
	}
	return b
}

func (b *ApplicationCriteriaBuilder) Build() Criteria {
	return &CompositeCriteria{criteriaList: b.criteriaList}
}
