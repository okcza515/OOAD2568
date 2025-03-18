package controller

import (
	"ModEd/common/model"

	"gorm.io/gorm"
)

type FacultyController struct {
	Connector *gorm.DB
}

func CreateFacultyController(connector *gorm.DB) *FacultyController {
    faculty := FacultyController{Connector: connector}
	connector.AutoMigrate(&model.Faculty{})
	return &faculty
}

func (faculty *FacultyController) GetAllFaculties() ([]*model.Faculty, error) {
	faculties := []*model.Faculty{}
	result := faculty.Connector.Find(&faculty)
	return faculties, result.Error
}