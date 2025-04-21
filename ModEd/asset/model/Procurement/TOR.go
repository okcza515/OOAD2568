// MEP-1014
package model

import (
	"time"

	"gorm.io/gorm"
)

type TOR struct {
	TORID               uint           `gorm:"primaryKey"`
	InstrumentRequestID uint           `gorm:"index"`
	Scope               string         `gorm:"type:text;not null"`
	Deliverables        string         `gorm:"type:text"`
	Timeline            string         `gorm:"type:text"`
	Committee           string         `gorm:"type:text"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	InstrumentRequest   InstrumentRequest
}
