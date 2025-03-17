package model

import (
	"time"
	"gorm.io/gorm"
)

type Assessment struct {
	gorm.Model
	ID          string    `json:"ID"`
	ProjectID   string    `json:"ProjectID"`
	AssessmentCriteria  AssessmentCriteria `json:"AssessmentCriteria"`
}


type AssessmentRepository interface {
	Create(Assessment) (Assessment, error)
	Update(Assessment) (Assessment, error)
	Delete(Assessment) error
}