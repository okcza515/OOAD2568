package model

import (
	"ModEd/core"
	"errors"

	"gorm.io/gorm"
)

type Faculty struct {
	core.BaseModel
	Name   string `gorm:"not null;unique" csv:"name" json:"name"`
	Budget int    `gorm:"default:0" csv:"budget" json:"budget"`
}

func (Faculty) TableName() string {
	return "faculties"
}

func (f Faculty) Validate() error {
	if f.Name == "" {
		return errors.New("faculty name is required")
	}
	if f.Budget < 0 {
		return errors.New("budget cannot be negative")
	}
	return nil
}

func SetFacultyBudget(db *gorm.DB, name string, newBudget int) error {
	return db.Model(&Faculty{}).
		Where("name = ?", name).
		Update("budget", newBudget).Error
}

func UpdateFacultyBudget(db *gorm.DB, name string, delta int) error {
	if delta >= 0 {
		return db.Model(&Faculty{}).
			Where("name = ?", name).
			Update("budget", gorm.Expr("budget + ?", delta)).Error
	}
	return db.Model(&Faculty{}).
		Where("name = ?", name).
		Where("budget + ? >= 0", delta).
		Update("budget", gorm.Expr("budget + ?", delta)).Error
}
