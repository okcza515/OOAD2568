package model

import (
	"time"

	"ModEd/core"
)

type AssessmentSubmission struct {
	core.BaseModel
	StudentCode string `gorm:"foreignKey:StudentCode;references:StudentCode"`
	Answers     PathFile
	Submitted   bool      `gorm:"default:false"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at" validate:"-"`
	Score       float64   `gorm:"default:0.0"`
	Feedback    string    `gorm:"type:varchar(255);"`
}
