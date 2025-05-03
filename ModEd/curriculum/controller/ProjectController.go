// MEP-1008

package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"
	projectModel "ModEd/project/model"

	"gorm.io/gorm"
)

type ProjectControllerService interface {
	UpdateProject(data *model.ProjectEvaluation) error
	GetProject(id uint, preloads ...string) (*model.ProjectEvaluation, error)
	DeleteProject(id uint) error
	ListProject(condition map[string]interface{}) ([]*model.ProjectEvaluation, error)

	CreateEvaluation(body *model.ProjectEvaluation, strategyType string, criteria []projectModel.AssessmentCriteria) error
	GetProjectByAdvisorID(advisorID uint) (*[]model.ProjectEvaluation, error)
	GetProjectByCommitteeID(committeeID uint) (*[]model.ProjectEvaluation, error)
}

type ProjectController struct {
	db   *gorm.DB
	core *core.BaseController[*model.ProjectEvaluation]
}

func CreateProjectController(db *gorm.DB) ProjectControllerService {
	return &ProjectController{
		db:   db,
		core: core.NewBaseController[*model.ProjectEvaluation](db),
	}
}

func (c *ProjectController) CreateEvaluation(evaluation *model.ProjectEvaluation, strategyType string, criteria []projectModel.AssessmentCriteria) error {
	ctx := model.NewEvaluationContext(strategyType)
	evaluation.Score, evaluation.Comment = ctx.Evaluate(evaluation, criteria)
	return c.core.Insert(evaluation)
}

func (c *ProjectController) GetProjectByAdvisorID(advisorID uint) (*[]model.ProjectEvaluation, error) {
	evaluations := new([]model.ProjectEvaluation)
	err := c.db.Where("advisor_id = ?", advisorID).Find(evaluations).Error
	return evaluations, err
}

func (c *ProjectController) GetProjectByCommitteeID(committeeID uint) (*[]model.ProjectEvaluation, error) {
	evaluations := new([]model.ProjectEvaluation)
	err := c.db.Where("committee_id = ?", committeeID).Find(evaluations).Error
	return evaluations, err
}

func (c *ProjectController) UpdateProject(data *model.ProjectEvaluation) error {
	return c.core.UpdateByID(data)
}

func (c *ProjectController) GetProject(id uint, preloads ...string) (*model.ProjectEvaluation, error) {
	return c.core.RetrieveByID(id, preloads...)
}

func (c *ProjectController) DeleteProject(id uint) error {
	return c.core.DeleteByID(id)
}

func (c *ProjectController) ListProject(condition map[string]interface{}) ([]*model.ProjectEvaluation, error) {
	return c.core.List(condition)
}
