package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeniorProject struct {
	gorm.Model
	SeniorProjectId uuid.UUID `gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	GroupName       string    `gorm:"not null"`

	GroupMembers  []GroupMember  `gorm:"foreignKey:SeniorProjectID"`
	Advisors      []Advisor      `gorm:"foreignKey:SeniorProjectID"`
	Committees    []Committee    `gorm:"foreignKey:SeniorProjectID"`
	Assignments   []Assignment   `gorm:"foreignKey:SeniorProjectID"`
	Presentations []Presentation `gorm:"foreignKey:SeniorProjectID"`
	Reports       []Report       `gorm:"foreignKey:SeniorProjectID"`

	Assessment Assessment `gorm:"foreignKey:SeniorProjectID"`
}
