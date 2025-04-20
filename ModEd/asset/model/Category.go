package model

// MEP-1012 Asset

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string  `gorm:"type:varchar(255);not null"`
	Description  *string `gorm:"type:varchar(255)"`
	UpdatedBy    *uint
	DeletedBy    *uint
}
