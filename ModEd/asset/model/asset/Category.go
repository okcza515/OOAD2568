package asset

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryID   uuid.UUID    `gorm:"type:text;primaryKey" json:"category_id" csv:"category_id"`
	CategoryName string       `gorm:"type:varchar(255);not null" json:"category_name" csv:"category_name"`
	Description  *string      `gorm:"type:varchar(255)" json:"description,omitempty" csv:"description,omitempty"`
	UpdatedBy    *uuid.UUID   `gorm:"type:text" json:"updated_by" csv:"updated_by"`
	DeletedBy    *uuid.UUID   `gorm:"type:text" json:"deleted_by,omitempty" csv:"deleted_by,omitempty"`
	Instrument   []Instrument `gorm:"foreignKey:CategoryID;references:CategoryID;constraint:OnUpdate:CASCADE;"`
	Supply       []Supply     `gorm:"foreignKey:CategoryID;references:CategoryID;constraint:OnUpdate:CASCADE;"`
}
