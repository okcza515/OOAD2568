package controller

import (
	"ModEd/common/model"

	"gorm.io/gorm"
)

type Instructor struct {
	Connector *gorm.DB
}

func CreateInstructorController(connector *gorm.DB) *Instructor {
	instructor := Instructor{Connector: connector}
	connector.AutoMigrate(&model.Instructor{})
	return &instructor
}

func (instructor Instructor) GetAll() ([]*model.Instructor, error) {
	instructors := []*model.Instructor{}
	result := instructor.Connector.Find(&instructors)
	return instructors, result.Error
}

func (instructor Instructor) GetByInstructorId(instructorId string) (*model.Instructor, error) {
	i := &model.Instructor{}
	result := instructor.Connector.Where("instructor_id = ?", instructorId).First(i)
	return i, result.Error
}

func (instructor Instructor) Create(i *model.Instructor) error {
	return instructor.Connector.Create(i).Error
}

func (instructor Instructor) Update(instructorId string, updatedData map[string]any) error {
	return instructor.Connector.Model(&model.Instructor{}).
		Where("instructor_id = ?", instructorId).
		Updates(updatedData).Error
}

func (instructor Instructor) Delete(instructorId string) error {
	return instructor.Connector.Where("instructor_id = ?", instructorId).Delete(&model.Instructor{}).Error
}
