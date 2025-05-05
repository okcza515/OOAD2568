// MEP-1009 Student Internship
package model

import "ModEd/core"

type Company struct {
	core.BaseModel
	CompanyName     string `gorm:"type:varchar(255)"`
	MentorFirstName string `gorm:"type:varchar(255)"`
	MentorLastName  string `gorm:"type:varchar(255)"`
	MentorEmail     string `gorm:"type:varchar(255)"`
	MentorPhone     string `gorm:"type:varchar(255)"`
}
