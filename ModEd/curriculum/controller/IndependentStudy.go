package controller

import (
	"ModEd/curriculum/model"

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
