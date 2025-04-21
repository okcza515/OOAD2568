package controller

import (
	model "ModEd/curriculum/model/instructor-workload"

	projectModel "ModEd/project/model"

	"gorm.io/gorm"
)

type ProjectControllerInterface interface {
	CreateEvaluation(body *model.ProjectEvaluation, strategyType string, criteria []projectModel.AssessmentCriteria) error
	UpdateEvaluation(evaluationID uint, body *model.ProjectEvaluation) error
	DeleteEvaluation(evaluationID uint) error
	GetAllEvaluations() (*[]model.ProjectEvaluation, error)
	GetEvaluationByID(evaluationID uint) (*model.ProjectEvaluation, error)
}

type ProjectController struct {
	db *gorm.DB
}

func NewProjectController(db *gorm.DB) ProjectControllerInterface {
	return &ProjectController{db: db}
}

func (c *ProjectController) CreateEvaluation(body *model.ProjectEvaluation, strategyType string, criteria []projectModel.AssessmentCriteria) error {
	ctx := model.NewEvaluationContext(strategyType)

	score := ctx.Evaluate(criteria)
	body.Score = score

	result := c.db.Create(body)
	return result.Error
}

func (c *ProjectController) GetAllEvaluations() (*[]model.ProjectEvaluation, error) {
	evaluations := new([]model.ProjectEvaluation)
	result := c.db.Find(&evaluations)
	return evaluations, result.Error
}

func (c *ProjectController) GetEvaluationByID(evaluationID uint) (*model.ProjectEvaluation, error) {
	evaluation := new(model.ProjectEvaluation)
	result := c.db.First(&evaluation, "ID = ?", evaluationID)
	return evaluation, result.Error
}

func (c *ProjectController) UpdateEvaluation(evaluationID uint, body *model.ProjectEvaluation) error {
	body.ID = evaluationID
	result := c.db.Updates(body)
	return result.Error
}

func (c *ProjectController) DeleteEvaluation(evaluationID uint) error {
	result := c.db.Model(&model.ProjectEvaluation{}).Where("ID = ?", evaluationID).Update("deleted_at", nil)
	return result.Error
}
