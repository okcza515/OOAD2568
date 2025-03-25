package model

import (
	"time"

	"gorm.io/gorm"

	"ModEd/core"
)

type Student struct {
	gorm.Model
	StudentCode string        `gorm:"not null;unique" csv:"student_code"`
	FirstName   string        `csv:"first_name"`
	LastName    string        `csv:"last_name"`
	Email       string        `csv:"email"`
	StartDate   time.Time     `csv:"start_date"`
	BirthDate   time.Time     `csv:"birth_date"`
	Program     ProgramType   `csv:"program"`
	Status      StudentStatus `csv:"status"`
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

core.RegisterModel("Student", NewStudent)