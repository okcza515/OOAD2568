package model

import (
	"gorm.io/gorm"
)

type SeniorProject struct {
	gorm.Model
	GroupName     string         `gorm:"not null"`
	Members       []GroupMember  `gorm:"foreignKey:SeniorProjectId"`
	Advisors      []Advisor      `gorm:"foreignKey:SeniorProjectId"`
	Committees    []Committee    `gorm:"foreignKey:SeniorProjectId"`
	Assignments   []Assignment   `gorm:"foreignKey:SeniorProjectId"`
	Presentations []Presentation `gorm:"foreignKey:SeniorProjectId"`
	Reports       []Report       `gorm:"foreignKey:SeniorProjectId"`
	Assessments   []Assessment   `gorm:"foreignKey:SeniorProjectId"`
}
