//MEP-1009 Student Internship
package controller

import (
    "fmt"
    "gorm.io/gorm"
)

type BaseScoreController[T any] struct {
    Connector *gorm.DB
}

func (bc *BaseScoreController[T]) UpdateScore(studentID string, scoreFields map[string]interface{}, getApplication func(*gorm.DB, string) (uint, error)) error {
    applicationID, err := getApplication(bc.Connector, studentID)
    if err != nil {
        return fmt.Errorf("failed to find application for student_code '%s': %w", studentID, err)
    }

		if err := bc.Connector.Model(new(T)).Where("id = ?", applicationID).Updates(scoreFields).Error; err != nil {
        return fmt.Errorf("failed to update scores for record id '%d': %w", applicationID, err)
    }

    return nil
}