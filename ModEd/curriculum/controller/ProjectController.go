// MEP-1008

package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"
	projectModel "ModEd/project/model"

	"gorm.io/gorm"
)

type ProjectControllerService interface {
	UpdateByID(data *model.ProjectEvaluation) error
	RetrieveByID(id uint, preloads ...string) (*model.ProjectEvaluation, error)
	DeleteByID(id uint) error
	List(condition map[string]interface{}) ([]*model.ProjectEvaluation, error)

	CreateEvaluation(body *model.ProjectEvaluation, strategyType string, criteria []projectModel.AssessmentCriteria) error
	GetProjectByAdvisorID(advisorID uint) (*[]model.ProjectEvaluation, error)
	GetProjectByCommitteeID(committeeID uint) (*[]model.ProjectEvaluation, error)
}

type ProjectController struct {
	*core.BaseController[*model.ProjectEvaluation]
	Connector *gorm.DB
}

func CreateProjectController(db *gorm.DB) *ProjectController {
	return &ProjectController{
		BaseController: core.NewBaseController[*model.ProjectEvaluation](db),
		Connector:      db,
	}
}

func (c *ProjectController) CreateEvaluation(evaluation *model.ProjectEvaluation, strategyType string, criteria []projectModel.AssessmentCriteria) error {
	ctx := model.NewEvaluationContext(strategyType)
	evaluation.Score = ctx.Evaluate(evaluation, criteria)

	return c.Insert(evaluation)
}

func (c *ProjectController) GetProjectByAdvisorID(advisorID uint) (*[]model.ProjectEvaluation, error) {
	evaluations := new([]model.ProjectEvaluation)
	err := c.Connector.Where("advisor_id = ?", advisorID).Find(evaluations).Error
	return evaluations, err
}

func (c *ProjectController) GetProjectByCommitteeID(committeeID uint) (*[]model.ProjectEvaluation, error) {
	evaluations := new([]model.ProjectEvaluation)
	err := c.Connector.Where("committee_id = ?", committeeID).Find(evaluations).Error
	return evaluations, err
}
