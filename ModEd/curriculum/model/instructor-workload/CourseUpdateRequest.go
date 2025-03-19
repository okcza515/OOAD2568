package model

import (
	commonModel "ModEd/common/model"

	"time"

	"gorm.io/gorm"
)

type CourseUpdateRequest struct {
	gorm.Model
	CourseId    int                         `gorm:"not null;index"`
	RequestType CourseUpdateRequestTypeEnum `gorm:"type:string;not null"`
	UpdatedBy   commonModel.Instructor      `gorm:"foreignKey:UpdatedByID"`
	UpdatedAt   time.Time                   `gorm:"not null;autoUpdateTime"`
}

type CourseNameUpdate struct {
	gorm.Model
	CourseId  int                    `gorm:"not null;index"`
	NewName   string                 `gorm:"type:string;not null"`
	UpdatedBy commonModel.Instructor `gorm:"foreignKey:UpdatedByID"`
}

type CoursePrerequisiteUpdate struct {
	gorm.Model
	CourseId  int                    `gorm:"not null;index"`
	UpdatedBy commonModel.Instructor `gorm:"foreignKey:UpdatedByID"`
}
