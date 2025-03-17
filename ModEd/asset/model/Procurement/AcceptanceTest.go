package model

import (
	"github.com/google/uuid"
)

type AcceptanceTest struct {
	AcceptanceTestID     uuid.UUID                  `gorm:"type:uuid;primaryKey"` // PK
	TORID                uuid.UUID                  `gorm:"type:uuid;not null"`   // for connection
	AcceptanceCriteriaID AcceptanceCriteria         `gorm:"foreignKey:AcceptanceCriteriaID"`
	WorkflowID           AcceptanceApprovalWorkflow `gorm:"foreignKey:AcceptanceApprovalWorkflowID"`
	Results              string                     `gorm:"type:text"`
}
