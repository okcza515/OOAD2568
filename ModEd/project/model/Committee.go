package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Committee struct {
	gorm.Model
	CommitteeId 		uuid.UUID		`gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	SeniorProjectId 	uuid.UUID		`gorm:"type:text;not null;index"`
	InstructorId    	uuid.UUID		`gorm:"type:text;not null;index"`
}