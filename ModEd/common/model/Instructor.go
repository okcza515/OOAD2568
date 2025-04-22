package model

import (
	"time"

	"gorm.io/gorm"
)

type Instructor struct {
	gorm.Model
	InstructorCode string     `gorm:"not null;unique" csv:"instructor_code" json:"instructor_code"`
	FirstName      string     `gorm:"not null" csv:"first_name" json:"first_name"`
	LastName       string     `gorm:"not null" csv:"last_name" json:"last_name"`
	Email          string     `gorm:"not null" csv:"email" json:"email"`
	StartDate      *time.Time `csv:"start_date" json:"start_date"`
	Department     string    `csv:"department" json:"department"`
}

func GetAllInstructors(db *gorm.DB) ([]*Instructor, error) {
	var instructors []*Instructor
	result := db.Find(&instructors)
	return instructors, result.Error
}

func GetInstructorByCode(db *gorm.DB, code string) (*Instructor, error) {
	var instructor Instructor
	result := db.Where("instructor_code = ?", code).First(&instructor)
	return &instructor, result.Error
}

func CreateInstructor(db *gorm.DB, instructor *Instructor) error {
	return db.Create(instructor).Error
}

func UpdateInstructorByCode(db *gorm.DB, code string, updated map[string]any) error {
	return db.Model(&Instructor{}).Where("instructor_code = ?", code).Updates(updated).Error
}

func DeleteInstructorByCode(db *gorm.DB, code string) error {
	return db.Where("instructor_code = ?", code).Delete(&Instructor{}).Error
}

func TruncateInstructors(db *gorm.DB) error {
	return db.Exec("DELETE FROM instructors").Error
}

func RegisterInstructors(db *gorm.DB, instructors []*Instructor) error {
	for _, i := range instructors {
		if err := db.Create(&i).Error; err != nil {
			return err
		}
	}
	return nil
}
