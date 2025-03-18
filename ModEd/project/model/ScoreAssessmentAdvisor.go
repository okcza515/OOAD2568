package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScoreAssessmentAdvisor struct {
	gorm.Model
	ScoreAssessmentAdvisorId uuid.UUID `gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	AssessmentId             uuid.UUID `gorm:"not null;index"`
	AdvisorId                uuid.UUID `gorm:"not null;index"`
	Score                    float64   `gorm:"not null"`

	Assessment *Assessment `gorm:"foreignKey:AssignmentId"`
	Advisor    *Advisor    `gorm:"foreignKey:AdvisorId"`
}
