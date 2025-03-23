package model

import (
	"time"

	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	CourseId     uint           `gorm:"primaryKey" csv:"course_id" json:"course_id"`
	Name         string         `gorm:"not null" csv:"name" json:"name"`
	Description  string         `gorm:"not null" csv:"description" json:"description"`
	CurriculumId uint           `gorm:"not null" csv:"curriculum_id" json:"curriculum_id"`
	Curriculum   Curriculum     `gorm:"foreignKey:CurriculumId;references:CurriculumId" csv:"-" json:"-"`
	Optional     bool           `gorm:"not null" csv:"optional" json:"optional"`
	CourseStatus CourseStatus   `gorm:"type:text;not null" csv:"course_status" json:"course_status"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" csv:"created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `csv:"-" json:"-"`
	ClassList    []Class        `gorm:"foreignKey:CourseId;references:CourseId" csv:"-" json:"-"`
	Prerequisite []Course       `gorm:"many2many:course_prerequisites;foreignKey:CourseId;joinForeignKey:CourseId;References:CourseId;joinReferences:PrerequisiteId" csv:"-" json:"-"`
}
