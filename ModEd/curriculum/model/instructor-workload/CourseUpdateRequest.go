package model

import (
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type CourseUpdateRequest struct {
	gorm.Model
	CourseId    int                         `gorm:"not null;index"`
	Course      model.Course                `gorm:"foreignKey:CourseId;references:CourseId"`
	RequestType CourseUpdateRequestTypeEnum `gorm:"type:string;not null"`
	Audit
}

type CourseNameUpdate struct {
	gorm.Model
	CourseId int          `gorm:"not null;index"`
	Course   model.Course `gorm:"foreignKey:CourseId;references:CourseId"`
	NewName  string       `gorm:"type:string;not null"`
	Audit
}

type CoursePrerequisiteUpdate struct {
	gorm.Model
	CourseId int          `gorm:"not null;index"`
	Course   model.Course `gorm:"foreignKey:CourseId;references:CourseId"`
	Audit
}
