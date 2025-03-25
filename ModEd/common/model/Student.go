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

func (student *Student) GetID() uint {
	return 0
}

func (student *Student) ToString() string {
	return ""
}

func NewStudent() *Student {
	return &Student{}
}

// core.RegisterModel("Student", NewStudent)
