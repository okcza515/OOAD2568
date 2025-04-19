// MEP-1014
package model

import (
	"time"

	"gorm.io/gorm"
)

type TOR struct {
	TORID         uint           `gorm:"primaryKey"`
	ItemRequestID uint           `gorm:"not null"`
	Scope         string         `gorm:"type:text;not null"`
	Deliverables  string         `gorm:"type:text"`
	Timeline      string         `gorm:"type:text"`
	Committee     string         `gorm:"type:text"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
