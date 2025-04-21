package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type AssessmentCriteriaController struct {
	*core.BaseController
	db *gorm.DB
}

func NewAssessmentCriteriaController(db *gorm.DB) *AssessmentCriteriaController {
	return &AssessmentCriteriaController{
		db:             db,
		BaseController: core.NewBaseController("assessmentCriterias", db),
	}
}

func (c *AssessmentCriteriaController) ListAllAssessmentCriterias() ([]model.AssessmentCriteria, error) {
	var assessmentCriterias []model.AssessmentCriteria
	result := c.db.Find(&assessmentCriterias)
	if result.Error != nil {
		return nil, result.Error
	}
	return assessmentCriterias, nil
}

func (c *AssessmentCriteriaController) RetrieveAssessmentCriteria(id uint) (*model.AssessmentCriteria, error) {
	var assessmentCriteria model.AssessmentCriteria
	if err := c.db.First(&assessmentCriteria, id).Error; err != nil {
		return nil, err
	}
	return &assessmentCriteria, nil
}

func (c *AssessmentCriteriaController) InsertAssessmentCriteria(criteriaName string) error {
	assessmentCriteria := model.AssessmentCriteria{
		CriteriaName: criteriaName,
	}
	return c.Insert(&assessmentCriteria)
}

func (c *AssessmentCriteriaController) UpdateAssessmentCriteria(id uint, assessmentCriteria *model.AssessmentCriteria) error {
	return c.UpdateByID(assessmentCriteria)
}

func (c *AssessmentCriteriaController) DeleteAssessmentCriteria(id uint) error {
	// Check if the criteria is used in any assessment link
	var links []model.AssessmentCriteriaLink
	err := c.db.Where("assessment_criteria_id = ?", id).Find(&links).Error
	if err != nil {
		return fmt.Errorf("error checking criteria usage: %v", err)
	}

	if len(links) > 0 {
		var projectIDs []string
		for _, link := range links {
			var assessment model.Assessment
			err := c.db.First(&assessment, link.AssessmentId).Error
			if err == nil {
				projectIDs = append(projectIDs, fmt.Sprintf("%d", assessment.SeniorProjectId))
			}
		}
		return fmt.Errorf("cannot delete criteria, used in SeniorProjectIDs: [%s]", strings.Join(projectIDs, ", "))
	}

	return c.db.Delete(&model.AssessmentCriteria{}, id).Error
}
