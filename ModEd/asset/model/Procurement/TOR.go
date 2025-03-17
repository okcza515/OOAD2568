package model

import (
	"github.com/google/uuid"
)

type TOR struct {
	TORID		  uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProcurementID uuid.UUID `gorm:"type:uuid;not null"`
	Scope         string    `gorm:"type:text"`
	Deliverables  string    `gorm:"type:text"`
	Timeline      string    `gorm:"type:text"`
}