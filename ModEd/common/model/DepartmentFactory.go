package model

import (
	// "fmt"

	"gorm.io/gorm"
)

type DepartmentFactory struct {
}

func NewDepartment(db *gorm.DB, dept Department) (*Department, error) {
	// var existing Department
	// if err := db.Where("name = ?", dept.Name).First(&existing).Error; err == nil {
	// 	return nil, fmt.Errorf("department with name '%s' already exists", dept.Name)
	// }

	// var faculty string
	// if err := db.Where("name = ?", dept.Faculty).First(&faculty).Error; err != nil {
	// 	return nil, fmt.Errorf("faculty '%s' not found: %w", dept.Faculty, err)
	// }

	return &Department{
		Name:    dept.Name,
		Faculty: dept.Faculty,
		Budget:  dept.Budget,
	}, nil
}