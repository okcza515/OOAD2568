package util

import (
	"encoding/csv"
	"fmt"
	"os"

	"ModEd/core"
	"ModEd/recruit/model"

	"gorm.io/gorm"
)

type CSVImporter struct {
	DB        *gorm.DB
	TableName string
}

// MapRowToModel แปลงข้อมูล CSV ตามประเภท table
func (ci *CSVImporter) MapRowToModel(row []string) (interface{}, error) {
	if len(row) < 1 {
		return nil, fmt.Errorf("invalid row data")
	}

	switch ci.TableName {
	case "application_rounds":
		return &model.ApplicationRound{
			RoundName: row[0],
		}, nil

	case "faculty":
		return &model.Faculty{
			Name: row[0],
		}, nil

	default:
		return nil, fmt.Errorf("unsupported table: %s", ci.TableName)
	}
}

func (ci *CSVImporter) ReadFromCSV(filePath string) error {
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

	ci.DB.Exec(fmt.Sprintf("DELETE FROM %s", ci.TableName))

	for _, row := range rows {
		modelInstance, err := ci.MapRowToModel(row)
		if err != nil {
			return fmt.Errorf("failed to map row: %w", err)
		}

		if err := ci.DB.Create(modelInstance).Error; err != nil {
			return fmt.Errorf("failed to insert data: %w", err)
		}
	}
	return nil
}

// ///////

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
