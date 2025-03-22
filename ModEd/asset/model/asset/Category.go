package asset

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string  `gorm:"type:varchar(255);not null"`
	Description  *string `gorm:"type:varchar(255)"`
	UpdatedBy    *uint
	DeletedBy    *uint
	Instrument   []Instrument `gorm:"foreignKey:CategoryID;references:CategoryID;constraint:OnUpdate:CASCADE;"`
	Supply       []Supply     `gorm:"foreignKey:CategoryID;references:CategoryID;constraint:OnUpdate:CASCADE;"`
}
