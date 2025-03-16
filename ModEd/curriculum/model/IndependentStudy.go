package model

import (
	"gorm.io/gorm"
	"time"
	"github.com/google/uuid"
)

type IndependentStudy struct{
    gorm.Model
    IndependentStudyId           uuid.UUID
    IndependentStudyTopic        string
    IndependentStudyContent      string
    EvaluationFromAdvisor        int
    AssignmentDate               time.Time
    TurninDate                   time.Time
    DueDate                      time.Time
}