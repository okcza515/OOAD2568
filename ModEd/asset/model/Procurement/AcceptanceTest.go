package model

import (
	"github.com/google/uuid"
)

type AcceptanceTest struct {
	AcceptanceTestID     uuid.UUID `gorm:"type:uuid;primaryKey"` // PK
	ProcurementID        uuid.UUID `gorm:"foreignKey:ProcurementID"`
	TORID                uuid.UUID `gorm:"type:uuid;not null"` // for connection
	AcceptanceCriteriaID uuid.UUID `gorm:"foreignKey:AcceptanceCriteriaID"`
	WorkflowID           uuid.UUID `gorm:"foreignKey:AcceptanceApprovalWorkflowID"`
	Results              string    `gorm:"type:text"`
}
