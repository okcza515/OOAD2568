package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type AssessmentCriteriaController struct {
	*core.BaseController[*model.AssessmentCriteria]
	DB *gorm.DB
}

func NewAssessmentCriteriaController(db *gorm.DB) *AssessmentCriteriaController {
	return &AssessmentCriteriaController{
		BaseController: core.NewBaseController[*model.AssessmentCriteria](db),
		DB:             db,
	}
}

func (c *AssessmentCriteriaController) ListAllAssessmentCriterias() ([]*model.AssessmentCriteria, error) {
	assessmentCriterias, err := c.List(map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return assessmentCriterias, nil
}

func (c *AssessmentCriteriaController) RetrieveAssessmentCriteria(id uint) (*model.AssessmentCriteria, error) {
	assessmentCriteria, err := c.RetrieveByID(id)
	if err != nil {
		return nil, err
	}
	return assessmentCriteria, nil
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
	var existinglinks []model.AssessmentCriteriaLink
	err := c.DB.Where("assessment_criteria_id = ?", id).Find(&existinglinks).Error
	if err != nil {
		return fmt.Errorf("error checking criteria usage: %v", err)
	}

	if len(existinglinks) > 0 {
		var projectIDs []string
		for _, link := range existinglinks {
			var assessment model.Assessment
			err := c.DB.First(&assessment, link.AssessmentId).Error
			if err == nil {
				projectIDs = append(projectIDs, fmt.Sprintf("%d", assessment.SeniorProjectId))
			}
		}
		return fmt.Errorf("cannot delete criteria, used in SeniorProjectIDs: [%s]", strings.Join(projectIDs, ", "))
	}

	return c.DeleteByID(id)
}
