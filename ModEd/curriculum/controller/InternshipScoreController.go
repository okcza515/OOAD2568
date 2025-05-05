//MEP-1009 Student Internship
package controller

import (
    "fmt"
    "gorm.io/gorm"
)

type BaseScoreController[T any] struct {
    Connector *gorm.DB
}

func (bsc *BaseScoreController[T]) UpdateScoreByID(id uint, fields map[string]interface{}) error {
    result := bsc.Connector.Model(new(T)).Where("id = ?", id).Updates(fields)
    if result.Error != nil {
        return fmt.Errorf("failed to update entity with ID %d: %w", id, result.Error)
    }
    if result.RowsAffected == 0 {
        return fmt.Errorf("no entity found with ID %d", id)
    }
    return nil
}