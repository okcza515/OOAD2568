// MEP-1007
package model

import (
	commonModel "ModEd/common/model"

	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	QuestionID uint                `gorm:"not null"`
	Question   Question            `gorm:"foreignKey:QuestionID"`
	StudentID  uint                `gorm:"not null"`
	Student    commonModel.Student `gorm:"foreignKey:StudentID"`
	Answer     string
}
