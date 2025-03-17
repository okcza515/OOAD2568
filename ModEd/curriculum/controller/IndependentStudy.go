package controller

import (
	"ModEd/curriculum/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IndependentStudyController struct {
	Connector *gorm.DB
}

func CreateIndependentStudyController(connector *gorm.DB) *IndependentStudyController {
	connector.AutoMigrate(&model.IndependentStudy{})
	return &IndependentStudyController{
		Connector: connector,
	}
}

func (repo IndependentStudyController) CreateIndependentStudy(independentStudies *[]model.IndependentStudy) {
	for _, independentStudy := range *independentStudies {
		repo.Connector.Create(independentStudy)
	}
}

func (repo IndependentStudyController) GetIndenpendentStudyByID(id uuid.UUID) (*model.IndependentStudy, error) {
	independentStudy := new(model.IndependentStudy)
	result := repo.Connector.First(&independentStudy, "IndependentStudyId = ?", id)
	return independentStudy, result.Error
}
