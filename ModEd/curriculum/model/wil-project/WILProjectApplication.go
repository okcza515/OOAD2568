package model

import "ModEd/common/model"

type WILProjectApplication struct {
	WILProjectApplicationId string              `json:"WILProjectApplicationId" gorm:"primaryKey;unique"`
	ProjectName             string              `json:"ProjectName"`
	ProjectDetail           string              `json:"ProjectDetail"`
	Company                 string              `json:"Company"`
	Mentor                  string              `json:"Mentor"`
	Students                []WILProjectMembers `json:"Students" gorm:"foreignKey:WILProjectApplicationId"`
	// Advisor                 Instructor      `json:"Advisor"` //TODO: Change this Instructor model to modelCommon.Instructor when its already implemented
	ApplicationStatus string `json:"ApplicationStatus"`
	TurninDate        string `json:"TurninDate"`
}

type WILProjectMembers struct {
	WILProjectApplicationId string        `json:"-" gorm:"primaryKey"`
	SID                     string        `json:"SID" gorm:"primaryKey"`
	Student                 model.Student `json:"Student" gorm:"foreignKey:SID;references:StudentID"`
}
