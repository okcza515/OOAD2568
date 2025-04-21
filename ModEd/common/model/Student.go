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
	Status      *StudentStatus `csv:"status" json:"status"`
}

func GetAllStudents(db *gorm.DB) ([]*Student, error) {
	var students []*Student
	result := db.Find(&students)
	return students, result.Error
}

func GetStudentByCode(db *gorm.DB, code string) (*Student, error) {
	var s Student
	result := db.Where("student_code = ?", code).First(&s)
	return &s, result.Error
}

func CreateStudent(db *gorm.DB, student *Student) error {
	return db.Create(student).Error
}

func UpdateStudentByCode(db *gorm.DB, code string, updated map[string]interface{}) error {
	return db.Model(&Student{}).Where("student_code = ?", code).Updates(updated).Error
}

func DeleteStudentByCode(db *gorm.DB, code string) error {
	return db.Where("student_code = ?", code).Delete(&Student{}).Error
}

func TruncateStudents(db *gorm.DB) error {
	return db.Exec("DELETE FROM students").Error
}

func RegisterStudents(db *gorm.DB, students []*Student) error {
	if err := TruncateStudents(db); err != nil {
		return err
	}
	for _, s := range students {
		if err := db.Create(s).Error; err != nil {
			return err
		}
	}
	return nil
}
