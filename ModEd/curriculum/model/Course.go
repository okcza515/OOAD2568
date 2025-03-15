package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	CourseId     uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" csv:"course_id" json:"course_id"`
	Name         string    `csv:"name" json:"name"`
	Optional     bool      `csv:"optional" json:"optional"`
	Prerequisite []Course  `csv:"-" json:"-"`
}
