package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Supply struct {
	gorm.Model
	SupplyID    uuid.UUID      `gorm:"type:text;primaryKey" json:"supply_id" csv:"supply_id"`
	SupplyLabel string         `gorm:"not null" json:"supply_label,omitempty" csv:"supply_label,omitempty"`
	Description string         `gorm:"type:text" json:"description,omitempty" csv:"description,omitempty"`
	RoomID      uuid.UUID      `gorm:"type:text;not null" json:"room_id,omitempty" csv:"room_id,omitempty"`
	Location    string         `gorm:"type:text" json:"location,omitempty" csv:"location,omitempty"`
	CategoryID  *uuid.UUID     `gorm:"type:text" json:"category_id,omitempty" csv:"category_id,omitempty"`
	Quantity    int            `gorm:"not null" json:"quantity,omitempty" csv:"quantity,omitempty"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" csv:"deleted_at,omitempty"`
}
