package model

import (
	"time"
	"gorm.io/gorm"
)

type Assessment struct {
	gorm.Model
	Id          uuid.UUID     `gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	ProjectId   uuid.UUID     `gorm:"type:text;not null;index"`
	
	AssessmentCriteria  AssessmentCriteria `gorm:"foreignKey:AssessmentId"`
	ScoreAssessment []Assessment `gorm:"foreignKey:AssessmentId"`
}