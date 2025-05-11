package controller

import (
	"encoding/json"
	"fmt"
	"os"
	"ModEd/asset/model"
	"gorm.io/gorm"
)

func ImportCriteriaFromJSON(db *gorm.DB, filename string) error {
	var criteriaList []model.AcceptanceCriteria

	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open JSON file: %v", err)
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&criteriaList); err != nil {
		return fmt.Errorf("failed to decode JSON: %v", err)
	}

	for _, criteria := range criteriaList {
		err := db.Create(&criteria).Error
		if err != nil {
			return fmt.Errorf("failed to insert criteria '%s': %v", criteria.CriteriaName, err)
		}
	}

	fmt.Println("Criteria imported successfully.")
	return nil
}
