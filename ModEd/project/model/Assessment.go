package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Assessment struct {
	gorm.Model
	AssessmentId    uuid.UUID `gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	SeniorProjectId uuid.UUID `gorm:"type:text;not null;index"`

	AssessmentCriteria      []*AssessmentCriteria       `gorm:"foreignKey:AssessmentCriteriaId"`
	ScoreAssessmentAdvisor  []*ScoreAssessmentAdvisor   `gorm:"foreignKey:AssessmentId"`
	ScoreAssessmentComittee []*ScoreAssessmentCommittee `gorm:"foreignKey:AssessmentId"`
	SeniorProject           *SeniorProject              `gorm:"foreignKey:SeniorProjectId"`
}
