package model

import (
	commonModel "ModEd/common/model"

	"time"

	"gorm.io/gorm"
)

type CourseUpdateRequest struct {
	gorm.Model
	CourseId    string                      `gorm:"not null"`
	RequestType CourseUpdateRequestTypeEnum `gorm:"not null"`
	UpdatedBy   commonModel.Instructor      `gorm:"not null"`
	UpdatedAt   time.Time                   `gorm:"not null"`

	UpdateStrategy CourseUpdateStrategy
}

type CourseNameUpdate struct {
	Connector     *gorm.DB
	CourseRequest *CourseUpdateRequest
}

type PrerequisiteUpdate struct {
	Connector     *gorm.DB
	CourseRequest *CourseUpdateRequest
}
