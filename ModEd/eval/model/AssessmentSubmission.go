package model

import (
	"time"

	"ModEd/core"
)

type AssessmentSubmission struct {
	core.BaseModel
	StudentCode string `gorm:"foreignKey:StudentCode;references:StudentCode"`
	FirstName   string `gorm:"foreignKey:StudentCode;references:FirstName"`
	LastName    string `gorm:"foreignKey:StudentCode;references:LastName"`
	Email       string `gorm:"foreignKey:StudentCode;references:Email"`
	Answers     PathFile
	Submitted   bool      `gorm:"default:false"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at" validate:"-"`
	Score       float64   `gorm:"default:0.0"`
	Feedback    string    `gorm:"type:varchar(255);"`
}
