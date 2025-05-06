//Chanawat Limpanatewin 65070503445
//MEP-1006

package controller

import (
	"ModEd/eval/model"
	"time"
)

type EvaluationController struct {
	evaluations []*model.Evaluation
	csvPath     string
}

func NewEvaluationController(evals []*model.Evaluation, csvPath string) *EvaluationController {
	return &EvaluationController{
		evaluations: evals,
		csvPath:     csvPath,
	}
}

func (ec *EvaluationController) EvaluateAssessment(studentCode, instructorCode string, assessmentID uint, assessmentType string, score uint, comment string) {
	for _, e := range ec.evaluations {
		if e.StudentCode == studentCode && e.AssessmentID == assessmentID && e.AssessmentType == assessmentType {
			e.Score = score
			e.Comment = comment
			e.EvaluatedAt = time.Now()
			model.SaveEvaluationsToCSV(ec.csvPath, ec.evaluations)
			return
		}
	}
	ec.evaluations = append(ec.evaluations, &model.Evaluation{
		StudentCode:    studentCode,
		InstructorCode: instructorCode,
		AssessmentID:   assessmentID,
		AssessmentType: assessmentType,
		Score:          score,
		Comment:        comment,
		EvaluatedAt:    time.Now(),
	})
	model.SaveEvaluationsToCSV(ec.csvPath, ec.evaluations)
}

func (ec *EvaluationController) ListEvaluations() []*model.Evaluation {
	return ec.evaluations
}
