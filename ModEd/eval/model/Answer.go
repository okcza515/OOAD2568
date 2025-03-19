package model

import (
	"gorm.io/gorm"
	commonModel "ModEd/common/model"
)

type Answer struct {
	gorm.Model
	ID 				uint					`gorm:"primaryKey"`
	Question		Question
	Student			commonModel.Student
	TheAnswer		string
}