package model

import (
	commonModel "ModEd/common/model"

	"gorm.io/gorm"
)

type WILProjectApplication struct {
	gorm.Model
	ProjectName       string                 `gorm:"not null"`
	ProjectDetail     string                 `gorm:"not null"`
	Semester          string                 `gorm:"not null"`
	Company           uint                   `gorm:"not null"`
	Mentor            string                 `gorm:"not null"`
	Students          []WILProjectMember     `gorm:"foreignKey:WILProjectApplicationId"`
	AdvisorId         uint                   `json:"AdvisorId"`
	Advisor           commonModel.Instructor `json:"Advisor", gorm:"foreignKey:AdvisorId;references:Id"`
	ApplicationStatus string                 `gorm:"not null"`
	TurninDate        string                 `gorm:"not null"`
}
