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

func CommonModelGetAll[T any](db *gorm.DB) ([]*T, error) {
	var models []*T
	result := db.Find(&models)
	return models, result.Error
}

func GetRecordByField[T CommonDataInterface](db *gorm.DB, field string, value interface{}) ([]*T, error) {
	var model []*T
	result := db.Where(fmt.Sprintf("%s = ?", field), value).Find(&model)
	return model, result.Error
}

// not finish
func DeleteRecordByField[T CommonDataInterface](db *gorm.DB, field string, value interface{}) error {
	var model []*T
	return db.Where(fmt.Sprintf("%s = ?", field), value).Delete(&model).Error
}

func TruncateModel(db *gorm.DB, model string) error {
	return db.Exec(fmt.Sprintf("DELETE FROM %s", model)).Error
}
