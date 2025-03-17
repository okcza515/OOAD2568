package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GroupMember struct {
	gorm.Model
	GroupMemberId   uuid.UUID `gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	SeniorProjectId uuid.UUID `gorm:"type:text;not null;index"`
	StudentId       uuid.UUID `gorm:"type:text;not null:index"`
}
