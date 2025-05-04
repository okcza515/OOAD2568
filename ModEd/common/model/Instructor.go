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
	Department     *string    `csv:"department" json:"department"`
}

func (Instructor) TableName() string {
	return "instructors"
}

func UpdateInstructorByCode(db *gorm.DB, code string, updated map[string]any) error {
	return db.Model(&Instructor{}).Where("instructor_code = ?", code).Updates(updated).Error
}

func DeleteInstructorByCode(db *gorm.DB, code string) error {
	return db.Where("instructor_code = ?", code).Delete(&Instructor{}).Error
}

func ManualAddInstructor(db *gorm.DB, instructor *Instructor) error {
	return db.Create(instructor).Error
}
