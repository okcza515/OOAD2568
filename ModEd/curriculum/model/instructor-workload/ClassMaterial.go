package model

import (
	"ModEd/curriculum/model"
)

type ClassMaterial struct {
	ClassMaterialId int         `gorm:"primaryKey;autoIncrement"`
	Class           model.Class `gorm:"foreignKey:ClassId;references:ClassId"`
	SourceUrl       string      `gorm:"type:text;not null"`
	CreatedAt       string      `gorm:"type:timestamp;not null"`
	UpdatedAt       string      `gorm:"type:timestamp;not null"`
}
