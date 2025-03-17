package model

import "github.com/google/uuid"

type WILProjectApplication struct {
	WILProjectApplicationId uuid.UUID          `gorm:"primaryKey;unique"`
	ProjectName             string             `gorm:"not null"`
	ProjectDetail           string             `gorm:"not null"`
	Company                 string             `gorm:"not null"`
	Mentor                  string             `gorm:"not null"`
	Students                []WILProjectMember `gorm:"foreignKey:WILProjectApplicationId"`
	// Advisor                 Instructor      `json:"Advisor"` //TODO: Change this Instructor model to modelCommon.Instructor when its already implemented
	ApplicationStatus string `gorm:"not null"`
	TurninDate        string `gorm:"not null"`
}
