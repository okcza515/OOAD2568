package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mentor struct {
	gorm.Model
	MentorID 					uuid.UUID
	MentorFirstName 	string
	MentorLastName		string
	MentorEmail 			string
	MentorPhone 			string
	Company 					Company	`gorm:"foreignKey:CompanyID;references:CompanyID"`
}