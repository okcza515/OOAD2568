//Chanawat Limpanatewin 65070503445
//MEP-1006

package controller

import (
	"ModEd/core"
	evalModel "ModEd/eval/model"
	"time"

	"gorm.io/gorm"
)

type EvaluationController struct {
	*core.BaseController[evalModel.Evaluation]
	db *gorm.DB
}

func NewEvaluationController(db *gorm.DB) *EvaluationController {
	return &EvaluationController{
		db:             db,
		BaseController: core.NewBaseController[evalModel.Evaluation](db),
	}
}

// CreateEvaluation creates a new evaluation
func (ec *EvaluationController) CreateEvaluation(studentCode, instructorCode string, assessmentId uint, score uint, comment string) error {
	newEvaluation := evalModel.Evaluation{
		StudentCode:    studentCode,
		InstructorCode: instructorCode,
		AssessmentId:   assessmentId,
		Score:          score,
		Comment:        comment,
		EvaluatedAt:    time.Now(),
	}
	return ec.Insert(newEvaluation)
}

// ViewAllEvaluations returns all evaluations with related data
func (ec *EvaluationController) ViewAllEvaluations() ([]evalModel.Evaluation, error) {
	return ec.List(nil, "Student", "Instructor", "Assessment")
}

// ViewEvaluationByID returns evaluation by student ID
func (ec *EvaluationController) ViewEvaluationByID(studentCode string) ([]evalModel.Evaluation, error) {
	condition := map[string]interface{}{
		"student_code": studentCode,
	}
	return ec.List(condition, "Student", "Instructor", "Assessment")
}

// UpdateEvaluation updates an existing evaluation
func (ec *EvaluationController) UpdateEvaluation(id uint, score uint, comment string) error {
	evaluation, err := ec.RetrieveByID(id, "Student", "Instructor", "Assessment")
	if err != nil {
		return err
	}

	evaluation.Score = score
	evaluation.Comment = comment
	evaluation.EvaluatedAt = time.Now()

	return ec.UpdateByID(evaluation)
}
