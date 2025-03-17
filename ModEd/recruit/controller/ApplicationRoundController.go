package controller

import (
	"ModEd/recruit/model"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/google/uuid"
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

func (ctrl *ApplicationRoundController) ReadApplicationRoundsFromCSV(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read CSV: %w", err)
	}

	ctrl.DB.Exec("DELETE FROM application_rounds")

	for _, row := range rows {
		if len(row) < 1 {
			continue
		}
		roundName := row[0]

		newRound := model.ApplicationRound{
			RoundID:   uuid.New(),
			RoundName: roundName,
		}
		if err := ctrl.DB.Create(&newRound).Error; err != nil {
			return fmt.Errorf("failed to insert application round: %w", err)
		}
	}
	return nil
}

func (arc *ApplicationRoundController) GetAllRounds() ([]*model.ApplicationRound, error) {
	var rounds []*model.ApplicationRound
	result := arc.DB.Find(&rounds)
	return rounds, result.Error
}
