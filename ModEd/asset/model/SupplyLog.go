package model

import (
	"gorm.io/gorm"
)

type SupplyLog struct {
	gorm.Model
	RefUserID   *uint
	StaffUserID uint            `gorm:"not null"`
	Action      SupplyLogAction `gorm:"not null"`
	SupplyID    uint            `gorm:"not null"`
	Description *string
	Quantity    uint `gorm:"not null;"`
}
