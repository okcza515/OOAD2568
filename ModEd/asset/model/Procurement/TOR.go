package model

import (
	"github.com/google/uuid"
)

type TOR struct {
	TORID         uuid.UUID          `gorm:"type:uuid;primaryKey"`
	ProcurementID uuid.UUID          `gorm:"type:uuid;not null"`
	Scope         string             `gorm:"type:text"`
	Deliverables  AcceptanceCriteria `gorm:"type:text"` //to-do: make it work properly.
	Timeline      string             `gorm:"type:text"`
}
