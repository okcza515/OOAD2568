package model

import (
	common "ModEd/common/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Advisor struct {
	gorm.Model
	AdvisorId			uuid.UUID	`gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	SeniorProjectId		uuid.UUID	`gorm:"type:text;not null;index"`
	InstructorId 		uuid.UUID	`gorm:"type:text;not null;index"`
	IsPrimary   		bool		`gorm:"not null"`

	SeniorProject	*SeniorProject		`gorm:"foreignKey:SeniorProjectId"`
	Instructor 		*common.Instructor	`gorm:"foreignKey:InstructorId"`
}
