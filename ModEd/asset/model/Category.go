package model

// MEP-1012 Asset

import (
	"ModEd/core"
)

type Category struct {
	core.BaseModel
	CategoryName string  `gorm:"type:varchar(255);not null"`
	Description  *string `gorm:"type:varchar(255)"`
	UpdatedBy    *uint
	DeletedBy    *uint
}
