package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"

	"gorm.io/gorm"
)

type AssessmentController struct {
	*core.BaseController
	db *gorm.DB
}

func NewAssessmentController(db *gorm.DB) *AssessmentController {
	return &AssessmentController{
		BaseController: core.NewBaseController("assessments", db),
		db:             db,
	}
}

// func (c *AssessmentController) ListAllAssessments() ([]core.RecordInterface, error) {
// 	return c.List(map[string]interface{}{})
// }

//#######################################
// for some reason, List and Retrieve function of BaseController doesn't know what Model it need to query
// even If I use AssessmentController that created from NewAssessmentController
// this problem occurs in other controllers as well when trying to retreive data with List or Retrieve of BaseController
// so I decided to use query of db itself instead
//#######################################

func (c *AssessmentController) ListAllAssessments() ([]model.Assessment, error) {
	var assessments []model.Assessment
	result := c.db.Find(&assessments)
	if result.Error != nil {
		return nil, result.Error
	}
	return assessments, nil
}

// func (c *AssessmentController) RetrieveAssessment(id uint) (*core.RecordInterface, error) {
// 	return c.RetrieveByID(id)
// }

func (c *AssessmentController) RetrieveAssessment(id uint) (*model.Assessment, error) {
	var assessment model.Assessment
	if err := c.db.First(&assessment, id).Error; err != nil {
		return nil, err
	}
	return &assessment, nil
}

func (c *AssessmentController) InsertAssessment(seniorProjectId uint) (*model.Assessment, error) {
	var existing []model.Assessment
	if err := c.db.Where("senior_project_id = ?", seniorProjectId).Find(&existing).Error; err != nil {
		return nil, fmt.Errorf("failed to check existing assessments: %w", err)
	}

	if len(existing) > 0 {
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
