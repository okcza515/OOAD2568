package model

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	CompanyName     string `gorm:"type:varchar(255)"`
	MentorFirstName string `gorm:"type:varchar(255)"`
	MentorLastName  string `gorm:"type:varchar(255)"`
	MentorEmail     string `gorm:"type:varchar(255)"`
	MentorPhone     string `gorm:"type:varchar(255)"`
}
