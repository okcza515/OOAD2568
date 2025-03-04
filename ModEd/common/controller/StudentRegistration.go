package controller

import (
	"ModEd/common/model"

	"gorm.io/gorm"
)

type StudentRegistration struct {
	Connector *gorm.DB
}

func CreateStudentRegistration(connector *gorm.DB) *StudentRegistration {
	registration := StudentRegistration{Connector: connector}
	connector.AutoMigrate(&model.Student{})
	return &registration
}

func (registration StudentRegistration) Register(students []*model.Student) {
	for _, student := range students {
		registration.Connector.Create(student)
	}
}

func (registration StudentRegistration) GetAll() ([]*model.Student, error) {
	students := []*model.Student{}
	result := registration.Connector.Find(&students)
	return students, result.Error
}
