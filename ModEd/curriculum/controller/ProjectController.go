// MEP-1008
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

	GetProjectByAdvisorID(advisorID uint) (*[]model.ProjectEvaluation, error)
	GetProjectByCommitteeID(committeeID uint) (*[]model.ProjectEvaluation, error)
}

type ProjectController struct {
	*core.BaseController[*model.ProjectEvaluation]
	Connector *gorm.DB
}

func CreateProjectController(db *gorm.DB) ProjectControllerService {
	return &ProjectController{
		BaseController: core.NewBaseController[*model.ProjectEvaluation](db),
		Connector:      db,
	}
}

func (c *ProjectController) CreateEvaluation(evaluation *model.ProjectEvaluation, strategyType string, criteria []projectModel.AssessmentCriteria) error {
	ctx := model.NewEvaluationContext(strategyType)

	score := ctx.Evaluate(evaluation, criteria)
	evaluation.Score = score

	result := c.Connector.Create(evaluation)
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

func (c *ProjectController) GetProjectByAdvisorID(advisorID uint) (*[]model.ProjectEvaluation, error) {
	evaluations := new([]model.ProjectEvaluation)
	result := c.Connector.Find(&evaluations, "advisor_id = ?", advisorID)
	return evaluations, result.Error
}

func (c *ProjectController) GetProjectByCommitteeID(committeeID uint) (*[]model.ProjectEvaluation, error) {
	evaluations := new([]model.ProjectEvaluation)
	result := c.Connector.Find(&evaluations, "committee_id = ?", committeeID)
	return evaluations, result.Error
}
