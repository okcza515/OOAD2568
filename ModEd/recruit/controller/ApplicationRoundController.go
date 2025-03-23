// MEP-1003 Student Recruitment
package controller

import (
	"ModEd/recruit/model"
	"ModEd/recruit/util"
	"log"

	"gorm.io/gorm"
)

type ApplicationRoundController struct {
	DB *gorm.DB
}

func CreateApplicationRoundController(db *gorm.DB) *ApplicationRoundController {
	err := db.AutoMigrate(&model.ApplicationRound{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v\n", err)
	}
	return &ApplicationRoundController{DB: db}

}

func (controller *ApplicationRoundController) CreateApplicationRound(round *model.ApplicationRound) error {
	result := controller.DB.Create(round)
	return result.Error
}

func (arc *ApplicationRoundController) ReadApplicationRoundsFromCSV(filePath string) error {
	importer := util.CSVImporter{
		DB:        arc.DB,
		TableName: "application_rounds",
	}

	return importer.ReadFromCSV(filePath)
}

func (arc *ApplicationRoundController) GetAllRounds() ([]*model.ApplicationRound, error) {
	var rounds []*model.ApplicationRound
	result := arc.DB.Find(&rounds)
	return rounds, result.Error
}
