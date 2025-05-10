package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"
	"fmt"

	"gorm.io/gorm"
)

type InternshipResultEvaluationController struct {
	*core.BaseController[model.InternshipResultEvaluation]
	Connector *gorm.DB
}

func NewInternshipResultEvaluationController(connector *gorm.DB) *InternshipResultEvaluationController {
	return &InternshipResultEvaluationController{
		Connector:      connector,
		BaseController: core.NewBaseController[model.InternshipResultEvaluation](connector),
	}
}

func (c *InternshipResultEvaluationController) Create(evaluation *model.InternshipResultEvaluation) error {
	if err := c.Connector.Create(evaluation).Error; err != nil {
		return fmt.Errorf("failed to create InternshipResultEvaluation: %w", err)
	}
	return nil
}

func (c *InternshipResultEvaluationController) RetrieveByID(id uint) (*model.InternshipResultEvaluation, error) {
	var evaluation model.InternshipResultEvaluation
	if err := c.Connector.First(&evaluation, id).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve InternshipResultEvaluation with ID %d: %w", id, err)
	}
	return &evaluation, nil
}

func (c *InternshipResultEvaluationController) Update(evaluation *model.InternshipResultEvaluation) error {
	if err := c.Connector.Save(evaluation).Error; err != nil {
		return fmt.Errorf("failed to update InternshipResultEvaluation with ID %d: %w", evaluation.ID, err)
	}
	return nil
}

func (c *InternshipResultEvaluationController) DeleteByID(id uint) error {
	if err := c.Connector.Delete(&model.InternshipResultEvaluation{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete InternshipResultEvaluation with ID %d: %w", id, err)
	}
	return nil
}

func (c *InternshipResultEvaluationController) ListAll() ([]model.InternshipResultEvaluation, error) {
	var evaluations []model.InternshipResultEvaluation
	if err := c.Connector.Find(&evaluations).Error; err != nil {
		return nil, fmt.Errorf("failed to list InternshipResultEvaluation records: %w", err)
	}
	return evaluations, nil
}
