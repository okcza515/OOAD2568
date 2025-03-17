package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ScoreAssignmentAdvisor struct {
	gorm.Model
	ScoreAssignmentAdvisorId 	uuid.UUID 		`gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	AssignmentId 				uuid.UUID    	`gorm:"not null;index"`
	AdvisorId    				uuid.UUID    	`gorm:"not null;index"`
	Score        				float64 		`gorm:"not null"`

	Assignment   				*Assignment 	`gorm:"foreignKey:AssignmentId"`
	Advisor 					*Advisor 		`gorm:"foreignKey:AdvisorId"`
}
