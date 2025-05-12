// MEP-1009 Student Internship
package model

import "ModEd/core"

type InternshipMentor struct {
	core.BaseModel
	MentorFirstName string  `gorm:"type:varchar(255)"`
	MentorLastName  string  `gorm:"type:varchar(255)"`
	MentorEmail     string  `gorm:"type:varchar(255)"`
	MentorPhone     string  `gorm:"type:varchar(255)"`
	CompanyId       uint    `gorm:"not null"`
	Company         Company `gorm:"foreignKey:CompanyId;references:ID"`
}
