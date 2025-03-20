// MEP-1014
package model

import (
	"github.com/google/uuid"
)

// refer from TOR
type AcceptanceCriteria struct {
	AcceptanceCriteriaID uuid.UUID `gorm:"type:uuid;primaryKey"` // PK
	CriteriaName         string    `gorm:"type:varchar(255);not null"`
	Description          string    `gorm:"type:text"`
}
