package model

import "gorm.io/gorm"

type ScoreAssessment struct {
	gorm.Model
	Id           string `gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	Score        int    `gorm:"type:integer; not null; index"`
	AssessmentId string `gorm:"type:text;not null;index"`
}
