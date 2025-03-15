package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Supply struct {
	SupplyID    uuid.UUID      `gorm:"type:text;primaryKey" json:"supply_id" csv:"supply_id"`
	SupplyLabel string         `gorm:"not null" json:"supply_label" csv:"supply_label"`
	Description *string        `gorm:"type:text" json:"description" csv:"description"`
	RoomID      uuid.UUID      `gorm:"type:text;not null" json:"room_id" csv:"room_id"`
	Location    *string        `gorm:"type:text" json:"location" csv:"location"`
	CategoryID  *uuid.UUID     `gorm:"type:text" json:"category_id" csv:"category_id"`
	Quantity    int            `gorm:"not null" json:"quantity" csv:"quantity"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at" csv:"deleted_at"`
}
