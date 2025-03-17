package model

import (
	"time"

	"github.com/google/uuid"
)

type AcceptanceTest struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"` // PK
	ProcurementID uuid.UUID `gorm:"type:uuid;not null"`
	Criteria      string    `gorm:"type:text"`
	Results       string    `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
