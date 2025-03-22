package asset

import (
	"gorm.io/gorm"
)

type SupplyLog struct {
	gorm.Model
	RefUserID   *string
	StaffUserID string          `gorm:"not null"`
	Action      SupplyLogAction `gorm:"not null"`
	SupplyID    string          `gorm:"not null"`
	Description *string
	Quantity    int `gorm:"not null;"`
}
