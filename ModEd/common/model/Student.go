package model

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	SID       string        `csv:"SID" gorm:"primarykey;size:16"`
	FirstName string        `csv:"FirstName"`
	LastName  string        `csv:"LastName"`
	Email     string        `csv:"Email"`
	StartDate time.Time     `csv:"-"`
	BirthDate time.Time     `csv:"-"`
	Program   ProgramType   `csv:"-"`
	Status    StudentStatus `csv:"-"`
}
