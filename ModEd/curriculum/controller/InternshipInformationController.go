package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"
	"fmt"

	"gorm.io/gorm"
)

type InternshipInformationController struct {
	*core.BaseController[model.InternshipInformation]
	Connector *gorm.DB
}

func NewInternshipInformationController(connector *gorm.DB) *InternshipInformationController {
	return &InternshipInformationController{
		Connector:      connector,
		BaseController: core.NewBaseController[model.InternshipInformation](connector),
	}
}

func (c *InternshipInformationController) Create(info *model.InternshipInformation) error {
	if err := c.Connector.Create(info).Error; err != nil {
		return fmt.Errorf("failed to create InternshipInformation: %w", err)
	}
	return nil
}

func (c *InternshipInformationController) RetrieveByID(id uint) (*model.InternshipInformation, error) {
	var info model.InternshipInformation
	if err := c.Connector.First(&info, id).Error; err != nil {
		return nil, fmt.Errorf("failed to retrieve InternshipInformation with ID %d: %w", id, err)
	}
	return &info, nil
}

func (c *InternshipInformationController) Update(info *model.InternshipInformation) error {
	if err := c.Connector.Save(info).Error; err != nil {
		return fmt.Errorf("failed to update InternshipInformation with ID %d: %w", info.ID, err)
	}
	return nil
}

func (c *InternshipInformationController) DeleteByID(id uint) error {
	if err := c.Connector.Delete(&model.InternshipInformation{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete InternshipInformation with ID %d: %w", id, err)
	}
	return nil
}

func (c *InternshipInformationController) ListAll() ([]model.InternshipInformation, error) {
	var infos []model.InternshipInformation
	if err := c.Connector.Find(&infos).Error; err != nil {
		return nil, fmt.Errorf("failed to list InternshipInformation records: %w", err)
	}
	return infos, nil
}
