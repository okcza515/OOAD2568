package model

import (
	commonModel "ModEd/common/model"

	"gorm.io/gorm"
)

type InternStudent struct {
	gorm.Model
	InternStatus InternStatus        `gorm:"type:varchar(20)"`
	StudentCode  string              `gorm:"type:varchar(255);not null"`
	Student      commonModel.Student `gorm:"foreignKey:StudentCode;references:StudentCode"`
}
