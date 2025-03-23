package model

import (
	"time"
	"gorm.io/gorm"
)

type Instructor struct {
	gorm.Model
	FirstName    string      `gorm:"not null" csv:"first_name" json:"first_name"`
	LastName     string      `gorm:"not null" csv:"last_name" json:"last_name"`
	Email        string      `gorm:"not null" csv:"email" json:"email"`
	StartDate    *time.Time   `csv:"start_date" json:"start_date"`
	Department   *string  	 `csv:"department" json:"department"`
	//course
}