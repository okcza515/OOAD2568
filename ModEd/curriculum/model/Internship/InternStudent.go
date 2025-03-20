package model

import (
	CommonModel "ModEd/common/model"
	"gorm.io/gorm"
)

type InternStudent struct {
	gorm.Model
	InternID     int                 `gorm:"primaryKey autoIncrement:true"`
	InternStatus InternStatus        `gorm:"type:varchar(20)"`
	StudentCode  uint                `gorm:"index"`
	Student      CommonModel.Student `gorm:"foreignKey:StudentID"`
}
