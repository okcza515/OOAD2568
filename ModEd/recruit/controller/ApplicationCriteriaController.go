// MEP-1003 Student Recruitment
package controller

type ApplicationCriteriaController struct{}

func NewApplicationCriteriaController() *ApplicationCriteriaController {
	return &ApplicationCriteriaController{}
}

func (ctrl *ApplicationCriteriaController) BuildCriteriaForApplicant(
	roundName string,
	facultyName string,
	departmentName string,
) Criteria {
	builder := NewApplicationCriteriaBuilder()

	builder.
		AddRoundCriteria(roundName).
		AddFacultyCriteria(facultyName).
		AddDepartmentCriteria(departmentName)

	return builder.Build()
}
