//MEP-1014
package model

import (
	"github.com/google/uuid"
)

type ItemRequest struct {
	ItemRequestID          uuid.UUID   `gorm:"type:uuid;primaryKey"` // PK
	ItemRequestDetailID    []uuid.UUID `gorm:"foreignKey:ItemRequestDetailID"`
	ItemApprovalWorkflowID uuid.UUID   `gorm:"foreignKey:ItemApprovalWorkflowID"`
	ItemBudgetAllocationID uuid.UUID   `gorm:"foreignKey:ItemBudgetAllocationID"`
}
