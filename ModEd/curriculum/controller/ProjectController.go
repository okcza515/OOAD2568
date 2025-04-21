//MEP-1008
package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"
	projectModel "ModEd/project/model"

	"gorm.io/gorm"
)

type ProjectControllerService interface {
	CreateEvaluation(body *model.ProjectEvaluation, strategyType string, criteria []projectModel.AssessmentCriteria) error
	UpdateEvaluation(evaluationID uint, body *model.ProjectEvaluation) error
	DeleteEvaluation(evaluationID uint) error
	GetAllEvaluations() (*[]model.ProjectEvaluation, error)
	GetEvaluationByID(evaluationID uint) (*model.ProjectEvaluation, error)
}

type ProjectController struct {
	*core.BaseController
	Connector *gorm.DB
}

func NewProjectController(db *gorm.DB) ProjectControllerService {
	return &ProjectController{
		BaseController: core.NewBaseController("Project", db),
		Connector:      db,
	}
}

func (c *ProjectController) CreateEvaluation(body *model.ProjectEvaluation, strategyType string, criteria []projectModel.AssessmentCriteria) error {
	ctx := model.NewEvaluationContext(strategyType)

	score := ctx.Evaluate(criteria)
	body.Score = score

	result := c.Connector.Create(body)
	return result.Error
}

func (c *ProjectController) GetAllEvaluations() (*[]model.ProjectEvaluation, error) {
	evaluations := new([]model.ProjectEvaluation)
	result := c.Connector.Find(&evaluations)
	return evaluations, result.Error
}

func (c *ProjectController) GetEvaluationByID(evaluationID uint) (*model.ProjectEvaluation, error) {
	evaluation := new(model.ProjectEvaluation)
	result := c.Connector.First(&evaluation, "ID = ?", evaluationID)
	return evaluation, result.Error
}

func (c *ProjectController) UpdateEvaluation(evaluationID uint, body *model.ProjectEvaluation) error {
	body.ID = evaluationID
	result := c.Connector.Updates(body)
	return result.Error
}

func (c *ProjectController) DeleteEvaluation(evaluationID uint) error {
	result := c.Connector.Model(&model.ProjectEvaluation{}).Where("ID = ?", evaluationID).Update("deleted_at", nil)
	return result.Error
}
