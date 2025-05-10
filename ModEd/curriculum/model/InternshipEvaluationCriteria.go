// MEP-1009 Student Internship
package model

import (
	"ModEd/core"
)

type InternshipEvaluationCriteria struct {
	core.BaseModel

	Title                   string                `gorm:"type:varchar(255);not null"`
	Description             string                `gorm:"type:varchar(255);not null"`
	Score                   uint                  `gorm:"not null"`
	InternshipApplication   InternshipApplication `gorm:"foreignKey:InternshipApplicationId;references:Id"`
	InternshipApplicationId uint                  `gorm:"not null"`
}
