package model

import (
	commonModel "ModEd/common/model"

	"gorm.io/gorm"
)

type StudentRequest struct {
	gorm.Model
	StudentId   int                 `gorm:"not null;index"`
	RequestType string              `gorm:"type:varchar(20);not null"`
	Remark      string              `gorm:"type:text"`
	Status      string              `gorm:"type:varchar(20);not null"`
	Student     commonModel.Student `gorm:"foreignKey:StudentId;references:ID"`
}
