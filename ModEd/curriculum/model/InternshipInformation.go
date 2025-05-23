// MEP-1009 Student Internship
package model

import "ModEd/core"

type InternshipInformation struct {
    core.BaseModel

    StudentCode        string           `gorm:"type:varchar(255);not null"`
    Student            InternStudent    `gorm:"foreignKey:StudentCode;references:StudentCode"`
    CompanyId          uint             `gorm:"not null"`
    InternshipCompany  Company          `gorm:"foreignKey:CompanyId;references:ID"`
    InternshipMentorID uint             `gorm:"not null"`
    InternshipMentor   InternshipMentor `gorm:"foreignKey:InternshipMentorID;references:ID"`
}