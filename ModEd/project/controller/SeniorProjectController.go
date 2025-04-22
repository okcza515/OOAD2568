package controller

import (
	"ModEd/project/model"

	"gorm.io/gorm"
)

type SeniorProjectController struct {
	DB *gorm.DB
}

func NewSeniorProjectController(db *gorm.DB) *SeniorProjectController {
	return &SeniorProjectController{DB: db}
}

func (controller *SeniorProjectController) InsertSeniorProject(seniorProject model.SeniorProject) error {
	return controller.DB.Create(&seniorProject).Error
}
