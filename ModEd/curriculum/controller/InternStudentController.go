//MEP-1009 Student Internship
package controller

import (
	commonmodel "ModEd/common/model"
	"ModEd/core"
	model "ModEd/curriculum/model"
	util "ModEd/utils/deserializer"
	"fmt"

	"gorm.io/gorm"
)

type InternStudentController struct {
	*core.BaseController[model.InternStudent]
	Connector *gorm.DB
}

func NewInternStudentController(connector *gorm.DB) *InternStudentController {
	return &InternStudentController{
		Connector:      connector,
		BaseController: core.NewBaseController[model.InternStudent](connector),
	}
}

func (repo InternStudentController) RegisterInternStudentsFromFile(filePath string) error {

	deserializer, err := util.NewFileDeserializer(filePath)
	if err != nil {
		return fmt.Errorf("failed to create deserializer: %w", err)
	}

	var internStudents []model.InternStudent
	if err := deserializer.Deserialize(&internStudents); err != nil {
		return fmt.Errorf("failed to deserialize file: %w", err)
	}

	for _, internStudent := range internStudents {

		var existingStudent commonmodel.Student
		if err := repo.Connector.Where("student_code = ?", internStudent.StudentCode).First(&existingStudent).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				fmt.Printf("Warning: Student with code %s does not exist. Skipping...\n", internStudent.StudentCode)
				continue
			}
			return fmt.Errorf("failed to check student_code %s: %w", internStudent.StudentCode, err)
		}

		internStudent.InternStatus = model.NOT_STARTED

		if err := repo.Connector.Create(&internStudent).Error; err != nil {
			return fmt.Errorf("failed to insert InternStudent: %w", err)
		}
	}

	return nil
}
