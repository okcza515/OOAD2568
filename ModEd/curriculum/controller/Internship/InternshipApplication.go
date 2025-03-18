package controller

import (
	model "ModEd/curriculum/model/Internship"

	"gorm.io/gorm"
)

type InternshipApplicationController struct {
	Connector *gorm.DB
}

func CreateInternshipApplicationController(connector *gorm.DB) *InternshipApplicationController {
	intern := InternshipApplicationController{Connector: connector}
	connector.AutoMigrate(&model.InternStudent{}, &model.InternshipApplication{})

	return &intern
}

func (repo InternshipApplicationController) RegisterInternshipApplication(application *model.InternStudent) {
	repo.Connector.Create(application)
}
