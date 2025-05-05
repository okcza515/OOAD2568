package model

import (
	commonModel "ModEd/common/model"

	"gorm.io/gorm"
)

type StatusResult string

const (
	PENDING StatusResult = "Pending"
	SUCCESS StatusResult = "Success"
)

type Result struct {
	gorm.Model
	ID            uint                `gorm:"primaryKey"`
	ExaminationID uint                `gorm:"not null"`
	Examination   Examination         `gorm:"foreignKey:ExaminationID"`
	StudentID     uint                `gorm:"not null"`
	Student       commonModel.Student `gorm:"foreignKey:StudentID"`
	Status        StatusResult        
	Feedback      string              
	Score         float64                
}
