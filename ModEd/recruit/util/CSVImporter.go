// MEP-1003 Student Recruitment
package util

import (
	"ModEd/core"

	"gorm.io/gorm"
)

type CSVImporter struct {
	DB        *gorm.DB
	TableName string
}

func ReadOnlyFromCSVOrJSON[T any](filePath string) ([]T, error) {
	mapper, err := core.CreateMapper[T](filePath)
	if err != nil {
		return nil, err
	}

	ptrRecords := mapper.Deserialize()
	var records []T
	for _, ptr := range ptrRecords {
		records = append(records, *ptr)
	}

	return records, nil
}

func InsertFromCSVOrJSON[T any](filePath string, db *gorm.DB) ([]T, error) {
	records, err := ReadOnlyFromCSVOrJSON[T](filePath)
	if err != nil {
		return nil, err
	}

	if err := db.Create(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}
