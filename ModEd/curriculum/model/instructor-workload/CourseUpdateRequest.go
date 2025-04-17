package model

import (
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type CourseUpdateRequest struct {
	gorm.Model
	ActionTracker
	CourseId    int                         `gorm:"not null;index"`
	Course      model.Course                `gorm:"foreignKey:CourseId;references:CourseId"`
	RequestType CourseUpdateRequestTypeEnum `gorm:"type:string;not null"`
}

type CourseNameUpdate struct {
	gorm.Model
	ActionTracker
	CourseId int          `gorm:"not null;index"`
	Course   model.Course `gorm:"foreignKey:CourseId;references:CourseId"`
	NewName  string       `gorm:"type:string;not null"`
}

type CoursePrerequisiteUpdate struct {
	gorm.Model
	ActionTracker
	CourseId int          `gorm:"not null;index"`
	Course   model.Course `gorm:"foreignKey:CourseId;references:CourseId"`
}
