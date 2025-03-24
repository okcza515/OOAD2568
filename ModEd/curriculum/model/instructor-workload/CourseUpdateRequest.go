package model

import (
	"gorm.io/gorm"
)

type CourseUpdateRequest struct {
	gorm.Model
	CourseId    int                         `gorm:"not null;index"`
	RequestType CourseUpdateRequestTypeEnum `gorm:"type:string;not null"`
	Audit
}

type CourseNameUpdate struct {
	gorm.Model
	CourseId int    `gorm:"not null;index"`
	NewName  string `gorm:"type:string;not null"`
	Audit
}

type CoursePrerequisiteUpdate struct {
	gorm.Model
	CourseId int `gorm:"not null;index"`
	Audit
}
