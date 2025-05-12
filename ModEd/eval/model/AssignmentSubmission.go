package model

import (
	"time"

	"ModEd/core"
)

type AssignmentSubmission struct {
	core.BaseModel
	AssignmentId uint      `gorm:"foreignKey:AssignmentId;references:AssignmentId"`
	StudentCode  string    `gorm:"column:student_code;type:varchar(255);"`
	Answer       string    `gorm:"column:answer;type:varchar(255);"`
	Submitted    bool      `gorm:"default:false"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at" validate:"-"`
}
