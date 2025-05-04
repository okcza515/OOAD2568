package model

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	StudentCode string         `gorm:"not null;unique" csv:"student_code" json:"student_code"`
	FirstName   string         `csv:"first_name" json:"first_name"`
	LastName    string         `csv:"last_name" json:"last_name"`
	Email       string         `csv:"email" json:"email"`
	StartDate   time.Time      `csv:"start_date" json:"start_date"`
	BirthDate   time.Time      `csv:"birth_date" json:"birth_date"`
	Program     ProgramType    `csv:"program" json:"program"`
	Department  string         `csv:"department" json:"department"`
	Status      *StudentStatus `csv:"status" json:"status"`
}

func (Student) TableName() string {
	return "students"
}

func UpdateStudentByCode(db *gorm.DB, code string, updated map[string]interface{}) error {
	return db.Model(&Student{}).Where("student_code = ?", code).Updates(updated).Error
}

func DeleteStudentByCode(db *gorm.DB, code string) error {
	return db.Where("student_code = ?", code).Delete(&Student{}).Error
}

func ManualAddStudent(db *gorm.DB, student *Student) error {
	return db.Create(student).Error
}