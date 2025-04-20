// MEP-1003 Student Recruitment
package controller

import criteria "ModEd/recruit/controller/Criteria"

type ApplicationCriteriaController struct{}

func NewApplicationCriteriaController() *ApplicationCriteriaController {
	return &ApplicationCriteriaController{}
}

func (ctrl *ApplicationCriteriaController) BuildCriteriaForApplicant(
	roundName string,
	facultyName string,
	departmentName string,
) criteria.Criteria {
	builder := criteria.NewApplicationCriteriaBuilder()

	builder.
		AddRoundCriteria(roundName).
		AddFacultyCriteria(facultyName).
		AddDepartmentCriteria(departmentName)

	return builder.Build()
}
