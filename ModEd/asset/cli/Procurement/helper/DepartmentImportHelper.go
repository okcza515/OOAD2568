package helper

import (
	"ModEd/common/model"
	"ModEd/utils/deserializer"
	"fmt"

	"gorm.io/gorm"
)

func SeedDepartmentsFromCSV(db *gorm.DB, path string) error {
	deser, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return fmt.Errorf("could not create deserializer: %w", err)
	}

	var departments []*model.Department
	if err := deser.Deserialize(&departments); err != nil {
		return fmt.Errorf("failed to deserialize CSV: %w", err)
	}

	for _, dept := range departments {
		if err := db.Create(dept).Error; err != nil {
			return fmt.Errorf("failed to insert department %s: %w", dept.Name, err)
		}
	}

	return nil
}
