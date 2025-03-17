package model

import "github.com/google/uuid"

type WILProjectApplication struct {
	WILProjectApplicationId uuid.UUID `gorm:"primaryKey;unique"`
	ProjectName             string
	ProjectDetail           string
	Company                 string
	Mentor                  string
	Students                []WILProjectMember `gorm:"foreignKey:WILProjectApplicationId"`
	// Advisor                 Instructor      `json:"Advisor"` //TODO: Change this Instructor model to modelCommon.Instructor when its already implemented
	ApplicationStatus string
	TurninDate        string
}
