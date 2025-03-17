package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	CourseId     uuid.UUID    `gorm:"type:uuid;primaryKey" csv:"course_id" json:"course_id"`
	Name         string       `csv:"name" json:"name"`
	Description  string       `json:"description"`
	CurriculumId uuid.UUID    `json:"curriculum_id"`
	Curriculum   Curriculum   `gorm:"foreignKey:CurriculumId;references:CurriculumId"`
	Optional     bool         `csv:"optional" json:"optional"`
	Prerequisite []Course     `csv:"-" json:"-"`
	CourseStatus CourseStatus `csv:"course_status" json:"course_status"`
	CreatedAt    time.Time
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
	ClassList    []Class
}

type CourseStatus string

const (
	ACTIVE   CourseStatus = "active"
	INACTIVE CourseStatus = "inactive"
)

var CourseStatusLabel = map[CourseStatus]string{
	ACTIVE:   "active",
	INACTIVE: "inactive",
}

func (c CourseStatus) String() string {
	return CourseStatusLabel[c]
}

func (c CourseStatus) IsValid() bool {
	_, ok := CourseStatusLabel[c]
	return ok
}
