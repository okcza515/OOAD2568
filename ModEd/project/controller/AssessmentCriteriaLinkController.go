package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"

	"gorm.io/gorm"
)

type AssessmentCriteriaLinkController struct {
	*core.BaseController[*model.AssessmentCriteriaLink]
	DB *gorm.DB
}

func NewAssessmentCriteriaLinkController(db *gorm.DB) *AssessmentCriteriaLinkController {
	return &AssessmentCriteriaLinkController{
		BaseController: core.NewBaseController[*model.AssessmentCriteriaLink](db),
		DB:             db,
	}
}

func (c *AssessmentCriteriaLinkController) ListAllAssessmentCriteriaLinks() ([]*model.AssessmentCriteriaLink, error) {
	assessmentCriteriaLinks, err := c.List(map[string]interface{}{})
	if assessmentCriteriaLinks != nil {
		return nil, err
	}
	return assessmentCriteriaLinks, nil
}

func (c *AssessmentCriteriaLinkController) ListProjectAssessmentCriteriaLinks(seniorProjectId uint) ([]*model.AssessmentCriteriaLink, error) {
	var assessment model.Assessment
	if err := c.DB.First(&assessment, "senior_project_id = ?", seniorProjectId).Error; err != nil {
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

// func (c *AssessmentCriteriaLinkController) UpdateAssessmentCriteriaLink(id uint, assessmentCriteriaLink *model.AssessmentCriteriaLink) error {
// 	return c.UpdateByID(assessmentCriteriaLink)
// }

func (c *AssessmentCriteriaLinkController) DeleteAssessmentCriteriaLink(assessmentID uint, criteriaID uint) error {
	return c.DB.
		Where("assessment_id = ? AND assessment_criteria_id = ?", assessmentID, criteriaID).
		Delete(&model.AssessmentCriteriaLink{}).
		Error
}
