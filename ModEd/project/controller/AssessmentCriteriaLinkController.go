package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"

	"gorm.io/gorm"
)

type AssessmentCriteriaLinkController struct {
	*core.BaseController[*model.AssessmentCriteriaLink]
	AssessmentController    *AssessmentController
	SeniorProjectController *core.BaseController[*model.SeniorProject]
	DB                      *gorm.DB
}

func NewAssessmentCriteriaLinkController(db *gorm.DB) *AssessmentCriteriaLinkController {
	return &AssessmentCriteriaLinkController{
		BaseController:          core.NewBaseController[*model.AssessmentCriteriaLink](db),
		AssessmentController:    NewAssessmentController(db),
		SeniorProjectController: core.NewBaseController[*model.SeniorProject](db),
		DB:                      db,
	}
}

func (c *AssessmentCriteriaLinkController) ListProjectAssessmentCriteriaLinks(seniorProjectId uint) ([]*model.AssessmentCriteriaLink, error) {
	//check if assessment exist, if not, auto create the assessment only once
	assessment, err := c.AssessmentController.RetrieveAssessmentBySeniorProjectId(seniorProjectId)
	if err != nil {
		return nil, err
	}

	assessmentCriteriaLink, err := c.List(map[string]interface{}{"assessment_id": assessment.ID})
	if err != nil {
		return nil, err
	}

	return assessmentCriteriaLink, nil
}

func (c *AssessmentCriteriaLinkController) RetrieveAssessmentCriteriaLink(assessmentId uint, assessmentCriteriaId uint) (*model.AssessmentCriteriaLink, error) {
	assessmentCriteriaLink, err := c.RetrieveByCondition(map[string]interface{}{"assessment_id": assessmentId, "assessment_criteria_id": assessmentCriteriaId})
	if err != nil {
		return nil, err
	}
	return assessmentCriteriaLink, nil
}

func (c *AssessmentCriteriaLinkController) InsertAssessmentCriteriaLink(assessmentId uint, assessmentCriteriaId uint) (*model.AssessmentCriteriaLink, error) {
	var assessment model.Assessment
	if err := c.DB.First(&assessment, assessmentId).Error; err != nil {
		return nil, fmt.Errorf("assessment does not exist: %w", err)
	}

	var assessmentCriteria model.AssessmentCriteria
	if err := c.DB.First(&assessmentCriteria, assessmentCriteriaId).Error; err != nil {
		return nil, fmt.Errorf("assessment criteria does not exist: %w", err)
	}

	assessmentCriteriaLink := model.AssessmentCriteriaLink{
		AssessmentId:         assessmentId,
		AssessmentCriteriaId: assessmentCriteriaId,
	}
	return &assessmentCriteriaLink, c.Insert(&assessmentCriteriaLink)
}

func (c *AssessmentCriteriaLinkController) DeleteAssessmentCriteriaLink(assessmentID uint, criteriaID uint) error {
	return c.DB.
		Where("assessment_id = ? AND assessment_criteria_id = ?", assessmentID, criteriaID).
		Delete(&model.AssessmentCriteriaLink{}).
		Error
}
