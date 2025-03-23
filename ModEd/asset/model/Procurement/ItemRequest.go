// MEP-1014
package model

import (
	"gorm.io/gorm"
)

type ItemRequest struct {
	ItemRequestID          uint           `gorm:"primaryKey"` // PK
	ItemRequestDetailID    uint           `gorm:"foreignKey:ItemRequestDetailID"`
	ItemApprovalID         uint           `gorm:"foreignKey:ItemApprovalID"`
	ItemBudgetAllocationID uint           `gorm:"foreignKey:ItemBudgetAllocationID"`
	DeletedAt              gorm.DeletedAt `gorm:"index"`
}
