package model

import "gorm.io/gorm"

type ScoreAssessment struct {
	gorm.Model
	Id           string `gorm:"ID"`
	Assessment Assessment `gorm:"Assessment"`
	Student      Student `gorm:"Student"`
	Score        int    `gorm:"Score"`
}

type ScoreAssessmentRepository interface {
	Create(score *ScoreAssessment) error
	Update(scoreAssessment *ScoreAssessment, score int) error
	Delete(scoreAssessment *ScoreAssessment) error
	GetByAssessmentId(assesmentId string) (*ScoreAssessment, error)
}