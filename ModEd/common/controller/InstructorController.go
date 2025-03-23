package controller

import (
	"ModEd/common/model"

	"gorm.io/gorm"
)

type InstructorController struct {
	Connector *gorm.DB
}

func CreateInstructorController(connector *gorm.DB) *InstructorController {
	instructor := InstructorController{Connector: connector}
	connector.AutoMigrate(&model.Instructor{})
	return &instructor
}

func (instructor InstructorController) GetAll() ([]*model.Instructor, error) {
	instructors := []*model.Instructor{}
	result := instructor.Connector.Find(&instructors)
	return instructors, result.Error
}

func (instructor InstructorController) GetByInstructorId(instructorId string) (*model.Instructor, error) {
	i := &model.Instructor{}
	result := instructor.Connector.Where("instructor_id = ?", instructorId).First(i)
	return i, result.Error
}

func (instructor InstructorController) Register(instructors *[]model.Instructor) {
	for _, i := range *instructors {
		instructor.Connector.Create(i)
	}
}

func (instructor InstructorController) Update(instructorId string, updatedData map[string]any) error {
	return instructor.Connector.Model(&model.Instructor{}).
		Where("instructor_id = ?", instructorId).
		Updates(updatedData).Error
}

func (instructor InstructorController) DeleteByInstructorId(instructorId string) error {
	return instructor.Connector.Where("instructor_id = ?", instructorId).Delete(&model.Instructor{}).Error
}
