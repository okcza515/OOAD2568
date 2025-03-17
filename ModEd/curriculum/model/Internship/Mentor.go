package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mentor struct {
	gorm.Model
	MentorID        uuid.UUID `gorm:"primaryKey"`
	MentorFirstName string    `gorm:"type:varchar(255)"`
	MentorLastName  string    `gorm:"type:varchar(255)"`
	MentorEmail     string    `gorm:"type:varchar(255)"`
	MentorPhone     string    `gorm:"type:varchar(255)"`
	Company         Company   `gorm:"foreignKey:CompanyID;references:CompanyID"`
}
