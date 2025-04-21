package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"

	"gorm.io/gorm"
)

type AssessmentCriteriaLinkController struct {
	*core.BaseController
	db *gorm.DB
}

func NewAssessmentCriteriaLinkController(db *gorm.DB) *AssessmentCriteriaLinkController {
	return &AssessmentCriteriaLinkController{
		db:             db,
		BaseController: core.NewBaseController("assessmentCriteriaLinks", db),
	}
}

func (c *AssessmentCriteriaLinkController) ListAllAssessmentCriteriaLinks() ([]model.AssessmentCriteriaLink, error) {
	var assessmentCriteriaLinks []model.AssessmentCriteriaLink
	result := c.db.Find(&assessmentCriteriaLinks)
	if result.Error != nil {
		return nil, result.Error
	}
	return assessmentCriteriaLinks, nil
}

func (c *AssessmentCriteriaLinkController) ListProjectAssessmentCriteriaLinks(seniorProjectId uint) ([]model.AssessmentCriteriaLink, error) {
	var assessment model.Assessment
	if err := c.db.First(&assessment, "senior_project_id = ?", seniorProjectId).Error; err != nil {
		return nil, err
	}
	var assessmentCriteriaLink []model.AssessmentCriteriaLink
	result := c.db.Find(&assessmentCriteriaLink, "assessment_id = ?", assessment.ID)
	if result.Error != nil {
		return nil, result.Error
	}
	return assessmentCriteriaLink, nil
}

func (c *AssessmentCriteriaLinkController) RetrieveAssessmentCriteriaLink(assessmentId uint, assessmentCriteriaId uint) (*model.AssessmentCriteriaLink, error) {
	var assessmentCriteriaLink model.AssessmentCriteriaLink
	if err := c.db.First(&assessmentCriteriaLink, "assessment_id = ? AND assessment_criteria_id = ?", assessmentId, assessmentCriteriaId).Error; err != nil {
		return nil, err
	}
	return &assessmentCriteriaLink, nil
}

func (c *AssessmentCriteriaLinkController) InsertAssessmentCriteriaLink(assessmentId uint, assessmentCriteriaId uint) (*model.AssessmentCriteriaLink, error) {
	var assessment model.Assessment
	if err := c.db.First(&assessment, assessmentId).Error; err != nil {
		return nil, fmt.Errorf("assessment does not exist: %w", err)
	}

	var assessmentCriteria model.AssessmentCriteria
	if err := c.db.First(&assessmentCriteria, assessmentCriteriaId).Error; err != nil {
		return nil, fmt.Errorf("assessment criteria does not exist: %w", err)
	}

	assessmentCriteriaLink := model.AssessmentCriteriaLink{
		AssessmentId:         assessmentId,
		AssessmentCriteriaId: assessmentCriteriaId,
	}
	return &assessmentCriteriaLink, c.Insert(&assessmentCriteriaLink)
}

func (c *AssessmentCriteriaLinkController) UpdateAssessmentCriteriaLink(id uint, assessmentCriteriaLink *model.AssessmentCriteriaLink) error {
	return c.UpdateByID(assessmentCriteriaLink)
}

func (c *AssessmentCriteriaLinkController) DeleteAssessmentCriteriaLink(assessmentID uint, criteriaID uint) error {
	return c.db.
		Where("assessment_id = ? AND assessment_criteria_id = ?", assessmentID, criteriaID).
		Delete(&model.AssessmentCriteriaLink{}).
		Error
}
