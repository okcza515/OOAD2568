package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"
	"fmt"

	"gorm.io/gorm"
)

type InternshipCriteriaController struct {
	*core.BaseController[model.InternshipCriteria]
	Connector *gorm.DB
}

func NewInternshipCriteriaController(connector *gorm.DB) *InternshipCriteriaController {
	return &InternshipCriteriaController{
		Connector:      connector,
		BaseController: core.NewBaseController[model.InternshipCriteria](connector),
	}
}

func (c *InternshipCriteriaController) Create(criteria *model.InternshipCriteria) error {
	if err := c.Connector.Create(criteria).Error; err != nil {
		return fmt.Errorf("failed to create InternshipCriteria: %w", err)
	}
	return nil
}

func (c *InternshipCriteriaController) RetrieveByID(id uint) (*model.InternshipCriteria, error) {
	var criteria model.InternshipCriteria
	if err := c.Connector.First(&criteria, id).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve InternshipCriteria with ID %d: %w", id, err)
	}
	return &criteria, nil
}

func (c *InternshipCriteriaController) Update(criteria *model.InternshipCriteria) error {
	if err := c.Connector.Save(criteria).Error; err != nil {
		return fmt.Errorf("failed to update InternshipCriteria with ID %d: %w", criteria.ID, err)
	}
	return nil
}

func (c *InternshipCriteriaController) DeleteByID(id uint) error {
	if err := c.Connector.Delete(&model.InternshipCriteria{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete InternshipCriteria with ID %d: %w", id, err)
	}
	return nil
}

func (c *InternshipCriteriaController) ListAll() ([]model.InternshipCriteria, error) {
	var criteriaList []model.InternshipCriteria
	if err := c.Connector.Find(&criteriaList).Error; err != nil {
		return nil, fmt.Errorf("failed to list InternshipCriteria records: %w", err)
	}
	return criteriaList, nil
}
