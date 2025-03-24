package model

import (
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type ClassMaterial struct {
	gorm.Model
	ClassId   uint        `gorm:"index"`
	Class     model.Class `gorm:"foreignKey:ClassId;references:ClassId"`
	SourceUrl string      `gorm:"type:text;not null"`
	Audit
}
