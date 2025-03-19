//MEP-1014
package model

import (
	"github.com/google/uuid"
)

type Procurement struct {
	ProcurementID                 uuid.UUID   `gorm:"type:uuid;primaryKey"` // PK
	TORcandidate                  []uuid.UUID `gorm:"foreignKey:TORID"`
	ItemRequestID                 uuid.UUID   `gorm:"foreignKey:ItemRequestID"`
	ProcurementApprovalWorkflowID uuid.UUID   `gorm:"type:uuid;ProcurementApprovalWorkflowID"`
}
