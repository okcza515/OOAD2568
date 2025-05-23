// MEP-1009 Student Internship
package model

import "ModEd/core"

type InternshipCriteria struct {
    core.BaseModel

    Title                   string                `gorm:"type:varchar(255);not null"`
    Description             string                `gorm:"type:varchar(255);not null"`
    Score                   uint                  `gorm:"not null"`
    InternshipApplicationId uint                  `gorm:"not null"`
    InternshipApplication   InternshipApplication `gorm:"foreignKey:InternshipApplicationId;references:ID"`
}