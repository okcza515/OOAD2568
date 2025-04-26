package controller

type ApplicationCriteriaDirector struct {
	builder *ApplicationCriteriaBuilder
}

func NewApplicationCriteriaDirector(builder *ApplicationCriteriaBuilder) *ApplicationCriteriaDirector {
	return &ApplicationCriteriaDirector{builder: builder}
}

func (d *ApplicationCriteriaDirector) BuildCriteria(roundName, facultyName, departmentName string) Criteria {
	return d.builder.
		AddRoundCriteria(roundName).
		AddFacultyCriteria(facultyName).
		AddDepartmentCriteria(departmentName).
		Build()
}
