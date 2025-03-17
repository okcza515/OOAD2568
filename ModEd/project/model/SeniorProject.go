package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeniorProject struct {
	gorm.Model
	SeniorProjectId uuid.UUID `gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	GroupName       string    `gorm:"not null"`

	GroupMembers  []*GroupMember  `gorm:"foreignKey:SeniorProjectId"`
	Advisors      []*Advisor      `gorm:"foreignKey:SeniorProjectId"`
	Committees    []*Committee    `gorm:"foreignKey:SeniorProjectId"`
	Assignments   []*Assignment   `gorm:"foreignKey:SeniorProjectId"`
	Presentations []*Presentation `gorm:"foreignKey:SeniorProjectId"`
	Reports       []*Report       `gorm:"foreignKey:SeniorProjectId"`

	Assessment *Assessment `gorm:"foreignKey:SeniorProjectId"`
}
