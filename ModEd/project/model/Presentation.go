package model

import (
	"time"

	"gorm.io/gorm"
)

type Presentation struct {
	gorm.Model
	SeniorProjectId  uint             `gorm:"type:text;not null;index"`
	PresentationType PresentationType `gorm:"type:varchar(50);not null"`
	Date             time.Time        `gorm:"type:date;not null"`
}
