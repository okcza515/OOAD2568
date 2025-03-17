package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Assignment struct {
	gorm.Model
	AssignmentId			uuid.UUID	`gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	SeniorProjectId 		uuid.UUID   `gorm:"type:text;not null;index"`
	Name  					string    	`gorm:"not null"`
	Description				string    	`gorm:"not null"`
	SubmissionDate			*time.Time 	`gorm:"type:date"`
	DueDate         		time.Time 	`gorm:"type:date;not null"`

	ScoreAssignmentAdvisor		[]*ScoreAssessmentAdvisor 	`gorm:"foreignKey:AssignmentId"`
	ScoreAssignmentCommittee	[]*ScoreAssessmentCommittee `gorm:"foreignKey:AssignmentId"`

	SeniorProject	*SeniorProject	`gorm:"foreignKey:AssignmentId"`		
	Progress		[]*Progress		`gorm:"foreignKey:AssignmentId"`
}