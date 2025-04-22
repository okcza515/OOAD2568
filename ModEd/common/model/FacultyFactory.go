package model

import (
	"fmt"

	"gorm.io/gorm"
)

func NewFaculty(db *gorm.DB, faculty Faculty) (*Faculty, error) {
	var existing Faculty
	if err := db.Where("name = ?", faculty.Name).First(&existing).Error; err == nil {
		return nil, fmt.Errorf("faculty '%s' already exists", faculty.Name)
	}

	return &Faculty{
		Name:   faculty.Name,
		Budget: faculty.Budget,
	}, nil
}
