package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssessmentCriteria struct {
	gorm.Model
	AssessmentCriteriaId uuid.UUID `gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	CriteriaName         string    `gorm:"not null"`
}
