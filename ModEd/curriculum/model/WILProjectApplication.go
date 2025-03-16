package model

import "gorm.io/gorm"

type WILProjectApplication struct {
	gorm.Model
	WILProjectApplicationId string `json:"WILProjectApplicationId" gorm:"primaryKey"`
	ProjectName             string `json:"ProjectName"`
	ProjectDetail           string `json:"ProjectDetail"`
	Company                 string `json:"Company"`
	Mentor                  string `json:"Mentor"`
	// Students                []model.Student `json:"Students" gorm:"foreignKey:SID"`
	// Advisor                 Instructor      `json:"Advisor"` //TODO: Change this Instructor model to modelCommon.Instructor when its already implemented
	ApplicationStatus string `json:"ApplicationStatus"`
	TurninDate        string `json:"TurninDate"`
}
