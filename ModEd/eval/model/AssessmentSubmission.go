package model

import (
	"time"

	"ModEd/core"
)

type AssessmentSubmission struct {
	core.BaseModel
	StudentCode string    `gorm:"foreignKey:StudentCode;references:StudentCode"`
	PdfPathFile string    `gorm:"type:varchar(255);"`
	Submitted   bool      `gorm:"default:false"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at" validate:"-"`
	Score       float64   `gorm:"default:0.0"`
	Feedback    string    `gorm:"type:varchar(255);"`
}
