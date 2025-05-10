package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"
	"fmt"

	"gorm.io/gorm"
)

type InternshipCriteriaController struct {
	*core.BaseController[model.InternshipEvaluationCriteria]
	Connector *gorm.DB
}

func NewInternshipCriteriaController(connector *gorm.DB) *InternshipCriteriaController {
	return &InternshipCriteriaController{
		Connector:      connector,
		BaseController: core.NewBaseController[model.InternshipEvaluationCriteria](connector),
	}
}

func (c *InternshipCriteriaController) Create(criteria *model.InternshipEvaluationCriteria) error {
	if err := c.Connector.Create(criteria).Error; err != nil {
		return fmt.Errorf("failed to create InternshipEvaluationCriteria: %w", err)
	}
	return nil
}

func (c *InternshipCriteriaController) RetrieveByID(id uint) (*model.InternshipEvaluationCriteria, error) {
	var criteria model.InternshipEvaluationCriteria
	if err := c.Connector.First(&criteria, id).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve InternshipEvaluationCriteria with ID %d: %w", id, err)
	}
	return &criteria, nil
}

func (c *InternshipCriteriaController) Update(criteria *model.InternshipEvaluationCriteria) error {
	if err := c.Connector.Save(criteria).Error; err != nil {
		return fmt.Errorf("failed to update InternshipEvaluationCriteria with ID %d: %w", criteria.ID, err)
	}
	return nil
}

func (c *InternshipCriteriaController) DeleteByID(id uint) error {
	if err := c.Connector.Delete(&model.InternshipEvaluationCriteria{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete InternshipEvaluationCriteria with ID %d: %w", id, err)
	}
	return nil
}

func (c *InternshipCriteriaController) ListAll() ([]model.InternshipEvaluationCriteria, error) {
	var criteriaList []model.InternshipEvaluationCriteria
	if err := c.Connector.Find(&criteriaList).Error; err != nil {
		return nil, fmt.Errorf("failed to list InternshipEvaluationCriteria records: %w", err)
	}
	return criteriaList, nil
}
