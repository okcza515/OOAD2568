package model

// MEP-1012 Asset

import (
	"ModEd/core"
)

type SupplyLog struct {
	core.BaseModel
	RefUserID   *uint
	StaffUserID uint            `gorm:"not null"`
	Action      SupplyLogAction `gorm:"not null"`
	SupplyID    uint            `gorm:"not null"`
	Description *string
	Quantity    uint `gorm:"not null;"`
}
