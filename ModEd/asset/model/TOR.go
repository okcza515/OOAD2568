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
	Quotations          []Quotation    `gorm:"foreignKey:TORID"`
	Status              TORStatus      `gorm:"type:varchar(50);default:'announced'"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	InstrumentRequest   InstrumentRequest
}
