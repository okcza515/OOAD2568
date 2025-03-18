package model

import (
	"ModEd/curriculum/model"

	"github.com/google/uuid"
)

type ClassMaterial struct {
	ClassMaterialId uuid.UUID   `gorm:"type:string;default:uuid_generate_v4();primaryKey"`
	Class           model.Class `gorm:"foreignKey:ClassId;references:ClassId"`
	SourceUrl       string      `gorm:"type:text;not null"`
	CreatedAt       string      `gorm:"type:timestamp;not null"`
	UpdatedAt       string      `gorm:"type:timestamp;not null"`
}
