package controller

import (
	commonmodel "ModEd/common/model"
	"ModEd/core"
	"ModEd/curriculum/model"
	"ModEd/utils/deserializer"
	"fmt"

	"gorm.io/gorm"
)

type ImportStrategy interface {
	Deserialize(filePath string) error
	SaveAll(db *gorm.DB) error
}

type CompanyImportStrategy struct {
	Companies []model.Company
}

func (s *CompanyImportStrategy) Deserialize(filePath string) error {
	des, err := deserializer.NewFileDeserializer(filePath)
	if err != nil {
		return err
	}
	return des.Deserialize(&s.Companies)
}

func (s *CompanyImportStrategy) SaveAll(db *gorm.DB) error {
	for _, company := range s.Companies {
		if err := db.Create(&company).Error; err != nil {
			return err
		}
	}
	return nil
}

type InternStudentImportStrategy struct {
	Students []model.InternStudent
}

func (s *InternStudentImportStrategy) Deserialize(filePath string) error {
	des, err := deserializer.NewFileDeserializer(filePath)
	if err != nil {
		return err
	}
	return des.Deserialize(&s.Students)
}

func (s *InternStudentImportStrategy) SaveAll(db *gorm.DB) error {
	for _, intern := range s.Students {
		var existing commonmodel.Student
		if err := db.Where("student_code = ?", intern.StudentCode).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Printf("Warning: Student code %s not found. Skipping.\n", intern.StudentCode)
				continue
			}
			return err
		}
		intern.InternStatus = model.NOT_STARTED
		if err := db.Create(&intern).Error; err != nil {
			return err
		}
	}
	return nil
}

type GenericImportController struct {
	*core.BaseController[model.InternStudent]
	Connector *gorm.DB
}

func CreateGenericImportController(connector *gorm.DB) *GenericImportController {
	return &GenericImportController{
		Connector:      connector,
		BaseController: core.NewBaseController[model.InternStudent](connector),
	}
}

func (c *GenericImportController) ImportWithStrategy(filePath string, s ImportStrategy) error {
	if err := s.Deserialize(filePath); err != nil {
		return fmt.Errorf("deserialization failed: %w", err)
	}
	if err := s.SaveAll(c.Connector); err != nil {
		return fmt.Errorf("saving failed: %w", err)
	}
	return nil
}

func (c *GenericImportController) ImportCompanies(filePath string) error {
	return c.ImportWithStrategy(filePath, &CompanyImportStrategy{})
}

func (c *GenericImportController) ImportInternStudents(filePath string) error {
	return c.ImportWithStrategy(filePath, &InternStudentImportStrategy{})
}
