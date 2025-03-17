package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IndependentStudy struct {
	gorm.Model
	IndependentStudyId      uuid.UUID  `gorm:"type:text;primarykey" json:"independent_study_id" csv:"independent_study_id"`
	IndependentStudyTopic   string     `gorm:"not null" json:"independent_study_topic" csv:"independent_study_topic"`
	IndependentStudyContent string     `gorm:"not null" json:"independent_study_content" csv:"independent_study_content"`
	EvaluationFromAdvisor   int        `gorm:"not null;default:0" json:"evaluation_from_advisor" csv:"evaluation_from_advisor"`
	AssignmentDate          time.Time  `gorm:"not null" json:"assign_date" csv:"assign_date"`
	TurnInDate              *time.Time `gorm:"default:null" json:"turn_in_date" csv:"turn_in_date"`
	DueDate                 time.Time  `gorm:"not null" json:"due_date" csv:"due_date"`
}
