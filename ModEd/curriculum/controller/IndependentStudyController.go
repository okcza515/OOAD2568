package controller

import (
	model "ModEd/curriculum/model"

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
	result := repo.Connector.First(&independentStudy, "independent_study_id = ?", id)
	return independentStudy, result.Error
}

func (repo IndependentStudyController) GetAllIndenpendentStudy() ([]*model.IndependentStudy, error) {
	independentStudy := []*model.IndependentStudy{}
	result := repo.Connector.Find(&independentStudy, "DeletedAt IS NULL")
	return independentStudy, result.Error
}

func (repo IndependentStudyController) UpdateIndenpendentStudy(updatedStudy *model.IndependentStudy) error {
	result := repo.Connector.Save(updatedStudy)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
