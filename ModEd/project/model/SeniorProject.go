package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SeniorProject struct {
	gorm.Model
	SeniorProjectId uuid.UUID `gorm:"type:text;uuid;primaryKey"`
	GroupName       string    `gorm:"not null"`
}
