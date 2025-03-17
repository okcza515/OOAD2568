package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)
type ScoreAssessment struct {
	gorm.Model
	ScoreAssessmentAdvisorId uuid.UUID `gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	AssessmentId uuid.UUID    `gorm:"not null;index"`
	ComitteeId    uuid.UUID    `gorm:"not null;index"`
	Score        float64 `gorm:"not null"`

	Assessment   *Assessment `gorm:"foreignKey:AssignmentId"`
	Comittee      Comittee     `gorm:"foreignKey:ComitteeId"`
}
