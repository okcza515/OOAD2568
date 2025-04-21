package model

// MEP-1012 Asset

import (
	"ModEd/core"
)

type Supply struct {
	core.BaseModel
	SupplyLabel string `gorm:"not null"`
	Description *string
	Location    *string
	CategoryID  *uint
	Quantity    int         `gorm:"not null"`
	SupplyLog   []SupplyLog `gorm:"foreignKey:ID;references:ID"`
}
