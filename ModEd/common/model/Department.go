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

func TruncateDepartments(db *gorm.DB) error {
	return db.Exec("DELETE FROM departments").Error
}


func RegisterDepartments(db *gorm.DB, departments []*Department) error {
	for _, dept := range departments {
		newDept, err := NewDepartment(db, *dept)
		if err != nil {
			return err
		}
		if err := db.Create(newDept).Error; err != nil {
			return err
		}
	}
	return nil
}
