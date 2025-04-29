package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"

	"gorm.io/gorm"
)

type AssessmentController struct {
	*core.BaseController[*model.Assessment]
	DB *gorm.DB
}

func NewAssessmentController(db *gorm.DB) *AssessmentController {
	return &AssessmentController{
		BaseController: core.NewBaseController[*model.Assessment](db),
		DB:             db,
	}
}

func (c *AssessmentController) ListAllAssessments() ([]*model.Assessment, error) {
	assessments, err := c.List(map[string]interface{}{})
	if assessments != nil {
		return nil, err
	}

	return assessments, nil
}

func (c *AssessmentController) RetrieveAssessment(id uint) (*model.Assessment, error) {
	assessment, err := c.RetrieveByID(id)
	if err != nil {
		return nil, err
	}

	return assessment, nil
}

func (c *AssessmentController) RetrieveAssessmentBySeniorProjectId(seniorProjectId uint) (*model.Assessment, error) {
	assessment, err := c.RetrieveByCondition(map[string]interface{}{"senior_project_id": seniorProjectId})
	if err != nil {
		return nil, err
	}
	return assessment, nil
}

func (c *AssessmentController) InsertAssessment(seniorProjectId uint) (*model.Assessment, error) {
	existingRecords, err := c.List(map[string]interface{}{"senior_project_id": seniorProjectId})
	if err != nil {
		return nil, fmt.Errorf("failed to check existing assessments: %w", err)
	}

	if len(existingRecords) > 0 {
		return nil, fmt.Errorf("assessments already exist for project %d", seniorProjectId)
	}

	assessment := model.Assessment{
		SeniorProjectId: seniorProjectId,
	}

	return &assessment, c.Insert(&assessment)
}

func (c *AssessmentController) UpdateAssessment(id uint, assessment *model.Assessment) error {
	return c.UpdateByID(assessment)
}

func (c *AssessmentController) DeleteAssessment(id uint) error {
	return c.DeleteByID(id)
}
