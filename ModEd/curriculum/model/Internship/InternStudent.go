package model

import (
	CommonModel "ModEd/common/model"

	"gorm.io/gorm"
)

type InternStudent struct {
	gorm.Model
	InternshipApplicationId uint                  `gorm:"primaryKey autoIncrement:true"`
	InternshipApplication   InternshipApplication `gorm:"foreignKey:InternshipApplicationId;references:ID"`
	InternStatus            InternStatus          `gorm:"type:varchar(20)"`
	StudentCode             string                `gorm:"not null"`
	Student                 CommonModel.Student   `gorm:"foreignKey:StudentCode;references:StudentCode"`
}
