package model

import (
	"fmt"

	"gorm.io/gorm"
)

func CommonRegister[T CommonDataInterface](db *gorm.DB, records []*T) error {
	for _, record := range records {
		if err := db.Create(record).Error; err != nil {
			return err
		}
	}
	return nil
}

func GetAllCommonModels[T any](db *gorm.DB) ([]*T, error) {
	var models []*T
	result := db.Find(&models)
	return models, result.Error
}

func GetRecordByField[T CommonDataInterface](db *gorm.DB, field string, value interface{}) ([]*T, error) {
	var model []*T
	result := db.Where(fmt.Sprintf("%s = ?", field), value).Find(&model)
	return model, result.Error
}

func TruncateModel(db *gorm.DB, model string) error {
	return db.Exec(fmt.Sprintf("DELETE FROM %s", model)).Error
}
