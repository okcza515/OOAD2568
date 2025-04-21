package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name        string        `gorm:"not null" csv:"name" json:"name"`
	Faculty     Faculty       `gorm:"foreignKey:ParentId" json:"parent"`
	Students    []Student     `gorm:"foreignKey:DepartmentId" json:"students"`
	Instructors *[]Instructor `gorm:"foreignKey:DepartmentId" json:"instructors"`
	Budget      int           `gorm:"default:0" csv:"budget" json:"budget"`
}

func GetAllDepartments(db *gorm.DB) ([]*Department, error) {
	var departments []*Department
	result := db.Find(&departments)
	return departments, result.Error
}

func GetDepartmentByName(db *gorm.DB, name string) (*Department, error) {
	var dept Department
	result := db.Where("name = ?", name).First(&dept)
	return &dept, result.Error
}

func CreateDepartment(db *gorm.DB, dept *Department) error {
	var existing Department
	if err := db.Where("name = ?", dept.Name).First(&existing).Error; err == nil {
		return fmt.Errorf("department with name '%s' already exists", dept.Name)
	}
	return db.Create(dept).Error
}

func SetDepartmentBudget(db *gorm.DB, name string, newBudget int) error {
	return db.Model(&Department{}).
		Where("name = ?", name).
		Update("budget", newBudget).Error
}

func UpdateDepartmentBudget(db *gorm.DB, name string, delta int) error {
	if delta >= 0 {
		return db.Model(&Department{}).
			Where("name = ?", name).
			Update("budget", gorm.Expr("budget + ?", delta)).Error
	}
	return db.Model(&Department{}).
		Where("name = ?", name).
		Where("budget + ? >= 0", delta).
		Update("budget", gorm.Expr("budget + ?", delta)).Error
}
