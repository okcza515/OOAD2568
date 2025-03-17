package model

type WILProjectApplication struct {
	WILProjectApplicationId string             `json:"WILProjectApplicationId" gorm:"primaryKey;unique"`
	ProjectName             string             `json:"ProjectName"`
	ProjectDetail           string             `json:"ProjectDetail"`
	Company                 string             `json:"Company"`
	Mentor                  string             `json:"Mentor"`
	Students                []WILProjectMember `json:"Students" gorm:"foreignKey:WILProjectApplicationId"`
	// Advisor                 Instructor      `json:"Advisor"` //TODO: Change this Instructor model to modelCommon.Instructor when its already implemented
	ApplicationStatus string `json:"ApplicationStatus"`
	TurninDate        string `json:"TurninDate"`
}
