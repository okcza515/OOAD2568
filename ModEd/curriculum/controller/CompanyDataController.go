//MEP-1009 Student Internship
package controller

import (
	"ModEd/core"
	model "ModEd/curriculum/model"
	"ModEd/utils/deserializer"
	"fmt"

	"gorm.io/gorm"
)

type CompanyDataController struct {
	*core.BaseController
	Connector *gorm.DB
}

func CreateCompanyDataController(connector *gorm.DB) *CompanyDataController {
	return &CompanyDataController{
		Connector:      connector,
		BaseController: core.NewBaseController("CompanyDataController", connector),
	}
}

func (c *CompanyDataController) ImportCompaniesFromCSV(filePath string) error {
	fileDeserializer, err := deserializer.NewFileDeserializer(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file deserializer: %w", err)
	}

	var companies []model.Company
	if err := fileDeserializer.Deserialize(&companies); err != nil {
		return fmt.Errorf("failed to deserialize CSV file: %w", err)
	}

	for _, company := range companies {
		if err := c.Connector.Create(&company).Error; err != nil {
			return fmt.Errorf("failed to insert company into database: %w", err)
		}
	}

	return nil
}
