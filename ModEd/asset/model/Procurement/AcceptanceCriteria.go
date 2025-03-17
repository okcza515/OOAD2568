package model

import (
	"time"

	"github.com/google/uuid"
)

type AcceptanceCriteria struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"` // PK
	CriteriaName string    `gorm:"type:varchar(255);not null"`
	Description  string    `gorm:"type:text"`

	CreatedAt time.Time
	UpdatedAt time.Time
}
