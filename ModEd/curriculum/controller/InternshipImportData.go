//MEP-1009 Student Internship
package controller

import (
	commonmodel "ModEd/common/model"
	"ModEd/core"
	"ModEd/curriculum/model"
	"ModEd/utils/deserializer"
	"fmt"

	"gorm.io/gorm"
)

type GenericImportController struct {
	*core.BaseController
	Connector *gorm.DB
}

func CreateGenericImportController(connector *gorm.DB) *GenericImportController {
	return &GenericImportController{
		Connector:      connector,
		BaseController: core.NewBaseController("GenericImportController", connector),
	}
}

func (c *GenericImportController) ImportDataFromFile(filePath string, target interface{}, deserializerFunc func(string) (*deserializer.FileDeserializer, error)) error {
	fileDeserializer, err := deserializerFunc(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file deserializer: %w", err)
	}

	if err := fileDeserializer.Deserialize(target); err != nil {
		return fmt.Errorf("failed to deserialize file: %w", err)
	}

	switch v := target.(type) {
	case *[]model.Company:
		for _, company := range *v {
			if err := c.Connector.Create(&company).Error; err != nil {
				return fmt.Errorf("failed to insert company into database: %w", err)
			}
		}
	case *[]model.InternStudent:
		for _, internStudent := range *v {
			var existingStudent commonmodel.Student
			if err := c.Connector.Where("student_code = ?", internStudent.StudentCode).First(&existingStudent).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					fmt.Printf("Warning: Student with code %s does not exist. Skipping...\n", internStudent.StudentCode)
					continue
				}
				return fmt.Errorf("failed to check student_code %s: %w", internStudent.StudentCode, err)
			}

			internStudent.InternStatus = model.NOT_STARTED

			if err := c.Connector.Create(&internStudent).Error; err != nil {
				return fmt.Errorf("failed to insert InternStudent: %w", err)
			}
		}
	default:
		return fmt.Errorf("unsupported target type: %T", v)
	}

	return nil
}

func (c *GenericImportController) ImportCompaniesFromCSV(filePath string) error {
	return c.ImportDataFromFile(filePath, &[]model.Company{}, deserializer.NewFileDeserializer)
}

func (c *GenericImportController) RegisterInternStudentsFromFile(filePath string) error {
	return c.ImportDataFromFile(filePath, &[]model.InternStudent{}, deserializer.NewFileDeserializer)
}