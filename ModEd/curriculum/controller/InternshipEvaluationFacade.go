package controller

type InternshipEvaluationFacade struct {
	InformationController      *InternshipInformationController
	CriteriaController         *InternshipCriteriaController
	ResultEvaluationController *InternshipResultEvaluationController
	AttendanceController       *InternshipAttendanceController
}

func NewInternshipEvaluationFacade(
	infoController *InternshipInformationController,
	criteriaController *InternshipCriteriaController,
	resultEvaluationController *InternshipResultEvaluationController,
	attendanceController *InternshipAttendanceController,
) *InternshipEvaluationFacade {
	return &InternshipEvaluationFacade{
		InformationController:      infoController,
		CriteriaController:         criteriaController,
		ResultEvaluationController: resultEvaluationController,
		AttendanceController:       attendanceController,
	}
}
