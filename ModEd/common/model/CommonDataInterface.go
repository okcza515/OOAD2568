package model

import (
	"fmt"

	"gorm.io/gorm"
)

type CommonDataInterface interface {
	TableName() string
}

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

func UpdateRecordByField[T CommonDataInterface](db *gorm.DB, field string, value interface{}, updatedData map[string]any, modelType CommonDataInterface) error {
	return db.Model(&modelType).Where(fmt.Sprintf("%s = ?", field), value).Updates(updatedData).Error
}

func DeleteRecordByField[T CommonDataInterface](db *gorm.DB, field string, value interface{}, modelType CommonDataInterface) error {
	return db.Where(fmt.Sprintf("%s = ?", field), value).Delete(&modelType).Error
}

func TruncateModel(db *gorm.DB, model string) error {
	return db.Exec(fmt.Sprintf("DELETE FROM %s", model)).Error
}