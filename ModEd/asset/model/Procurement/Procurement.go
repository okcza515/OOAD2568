package model

import (
	"github.com/google/uuid"
)

type Procurement struct {
	ID             uuid.UUID        `gorm:"type:uuid;primaryKey"` // PK
	RequestID      string           `gorm:"unique;not null"`
	Description    string           `gorm:"type:text"`
	RequestedItems []RequestedItem  `gorm:"foreignKey:ProcurementID"`
	Budget         BudgetAllocation `gorm:"foreignKey:ProcurementID"`
	Approval       ApprovalWorkflow `gorm:"foreignKey:ProcurementID"`
	Supplier       Supplier         `gorm:"foreignKey:SupplierID"`
	Status         string           `gorm:"type:varchar(50);not null"`
}
