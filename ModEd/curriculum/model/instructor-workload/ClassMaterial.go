package model

import (
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type ClassMaterial struct {
	gorm.Model
	ActionTracker
	Class    model.Class `gorm:"foreignKey:ClassID;references:ID"`
	FileName string      `gorm:"type:varchar(100);not null"`
	FilePath string      `gorm:"type:varchar(255);not null"`
}
