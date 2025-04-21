package model

import (
	commonModel "ModEd/common/model"

	"github.com/google/uuid"
)

type WILProjectApplication struct {
	WILProjectApplicationId uuid.UUID              `gorm:"primaryKey;unique"`
	ProjectName             string                 `gorm:"not null"`
	ProjectDetail           string                 `gorm:"not null"`
	Company                 string                 `gorm:"not null"`
	Mentor                  string                 `gorm:"not null"`
	Students                []WILProjectMember     `gorm:"foreignKey:WILProjectApplicationId"`
	Advisor                 commonModel.Instructor `json:"Advisor"`
	ApplicationStatus       string                 `gorm:"not null"`
	TurninDate              string                 `gorm:"not null"`
}
