package model

import (
	"gorm.io/gorm"
)

type Faculty struct {
	gorm.Model
	Name   string `gorm:"not null;unique" csv:"name" json:"name"`
	Budget int    `gorm:"default:0" csv:"budget" json:"budget"`
}

func GetAllFaculties(db *gorm.DB) ([]*Faculty, error) {
	var faculties []*Faculty
	result := db.Find(&faculties)
	return faculties, result.Error
}

func GetFacultyByName(db *gorm.DB, name string) (*Faculty, error) {
	var faculty Faculty
	result := db.Where("name = ?", name).First(&faculty)
	return &faculty, result.Error
}

func CreateFaculty(db *gorm.DB, faculty *Faculty) error {
	return db.Create(faculty).Error
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

func TruncateFaculties(db *gorm.DB) error {
	return db.Exec("DELETE FROM faculties").Error
}

func RegisterFaculties(db *gorm.DB, faculties []*Faculty) error {
	if err := TruncateFaculties(db); err != nil {
		return err
	}
	for _, faculty := range faculties {
		if err := db.Create(faculty).Error; err != nil {
			return err
		}
	}
	return nil
}
