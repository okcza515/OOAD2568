// MEP-1014
package model

import (
	"time"

	"gorm.io/gorm"
)

type TOR struct {
	TORID         uint           `gorm:"primaryKey"`
	ItemRequestID uint           `gorm:"not null"`                         // Foreign key to approved request
	Scope         string         `gorm:"type:text;not null"`               // What this TOR is about
	Deliverables  string         `gorm:"type:text"`                        // Human-readable deliverables
	Timeline      string         `gorm:"type:text"`                        // Timeline info
	Committee     string         `gorm:"type:text"`                        // Names of committee/approvers
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

