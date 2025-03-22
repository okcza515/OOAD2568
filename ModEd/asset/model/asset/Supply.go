package asset

import (
	"gorm.io/gorm"
)

type Supply struct {
	gorm.Model
	SupplyLabel string `gorm:"not null"`
	Description *string
	RoomID      uint `gorm:"not null"`
	Location    *string
	CategoryID  *uint
	Quantity    int         `gorm:"not null"`
	SupplyLog   []SupplyLog `gorm:"foreignKey:ID;references:SupplyID"`
}
