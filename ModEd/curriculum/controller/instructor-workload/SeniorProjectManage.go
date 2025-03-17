package controller

import (
	model "ModEd/curriculum/model/instructor-workload"
	projectModel "ModEd/project/model"

	"errors"

	"gorm.io/gorm"
)

type SeniorProjectEvaluateController struct {
	Connector *gorm.DB
}

func (s SeniorProjectEvaluateController) GetSeniorProjectUnderAdvisor(instructorId string) ([]*projectModel.SeniorProject, error) {
	projects := []*projectModel.SeniorProject{}
	result := s.Connector.Find(&projects, "advisor_id = ?", instructorId)
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (e *SeniorProjectEvaluateController) EvaluateTask(evaluation *model.ProjectEvaluation) (float64, string, error) {
	switch evaluation.AssignmentType {
	case "assignment":
		return e.EvaluateAssignment(evaluation)
	case "presentation":
		return e.EvaluatePresentation(evaluation)
	case "report":
		return e.EvaluateReport(evaluation)
	default:
		return 0, "", errors.New("Invalid assignment type")
	}
}

func (e *SeniorProjectEvaluateController) EvaluateAssignment(evaluation *model.ProjectEvaluation) (float64, string, error) {
	score := 10.0
	comment := "Good performance on the assignment."

	return score, comment, nil
}

func (e *SeniorProjectEvaluateController) EvaluatePresentation(evaluation *model.ProjectEvaluation) (float64, string, error) {
	score := 20.0
	comment := "Great presentation skills."

	return score, comment, nil
}

func (e *SeniorProjectEvaluateController) EvaluateReport(evaluation *model.ProjectEvaluation) (float64, string, error) {
	score := 30.0
	comment := "Excellent report quality."

	return score, comment, nil
}
