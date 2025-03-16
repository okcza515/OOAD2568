package model

import (
	"gorm.io/gorm"
)

type Advisor struct {
	gorm.Model
	SeniorProjectId		uint64	`gorm:"not null;index"`
	Name 				string	`gorm:"not null;size:255"`
	Email    			string	`gorm:"not null;size:255"`
	InstructorID 		string	`gorm:"not null:size:255"`
	IsPrimary   		bool	`gorm:"not null"`
}
