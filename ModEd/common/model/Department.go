package model

import (

	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name    string `gorm:"not null" csv:"name" json:"name"`
	Faculty string `gorm:"not null" csv:"faculty" json:"parent"`
	Budget  int    `gorm:"default:0" csv:"budget" json:"budget"`
}

func (Department) TableName() string {
	return "departments"
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

func GetDepartmentsByFaculty(db *gorm.DB, faculty string) ([]*Department, error) {
	var departments []*Department
	result := db.Where("faculty = ?", faculty).Find(&departments)
	return departments, result.Error
}
