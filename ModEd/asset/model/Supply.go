package model

import (
	"gorm.io/gorm"
)

type Supply struct {
	gorm.Model
	SupplyLabel string `gorm:"not null"`
	Description *string
	Location    *string
	CategoryID  *uint
	Quantity    int         `gorm:"not null"`
	SupplyLog   []SupplyLog `gorm:"foreignKey:ID;references:ID"`
}
