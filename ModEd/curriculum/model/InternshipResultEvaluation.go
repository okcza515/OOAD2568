// MEP-1009 Student Internship
package model

import "ModEd/core"

type InternshipResultEvaluation struct {
    core.BaseModel
    Comment                 string                `gorm:"type:varchar(255);not null"`
    Score                   uint                  `gorm:"not null"`
    InternshipInformationId uint                  `gorm:"not null"`
    InternshipInformation   InternshipInformation `gorm:"foreignKey:InternshipInformationId;references:ID"`
}