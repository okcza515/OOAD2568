package controller

import (
	model "ModEd/curriculum/model/instructor-workload"
	"errors"

	"gorm.io/gorm"
)

type AssignmentEvaluation struct{}
type PresentationEvaluation struct{}
type ReportEvaluation struct{}

type SeniorProjectEvaluateController struct {
	Connector *gorm.DB
}

func (e *SeniorProjectEvaluateController) EvaluateTask(evaluation *model.ProjectEvaluation) (float64, string, error) {
	var evaluator model.ProjectEvaluationStrategy
	switch evaluation.AssignmentType {
	case "assignment":
		evaluator = &AssignmentEvaluation{}
	case "presentation":
		evaluator = &PresentationEvaluation{}
	case "report":
		evaluator = &ReportEvaluation{}
	default:
		return 0, "", errors.New("Invalid assignment type")
	}

	evaluation.SetEvaluationStrategy(evaluator)

	return evaluation.ExecuteEvaluation()
}

func (a *AssignmentEvaluation) Evaluate(evaluation model.ProjectEvaluation) (float64, string, error) {
	score := 10.0
	comment := ""

	return score, comment, nil
}

func (a *PresentationEvaluation) Evaluate(evaluation model.ProjectEvaluation) (float64, string, error) {
	score := 20.0
	comment := ""

	return score, comment, nil
}

func (a *ReportEvaluation) Evaluate(evaluation model.ProjectEvaluation) (float64, string, error) {
	score := 30.0
	comment := ""

	return score, comment, nil
}
