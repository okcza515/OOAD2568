package model

import (
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type ClassMaterial struct {
	gorm.Model
	ClassId  uint        `gorm:"not null" json:"class_id"`
	Class    model.Class `gorm:"foreignKey:ClassId;references:ClassId"`
	FileName string      `gorm:"type:varchar(100);not null"`
	FilePath string      `gorm:"type:varchar(255);not null"`
}
