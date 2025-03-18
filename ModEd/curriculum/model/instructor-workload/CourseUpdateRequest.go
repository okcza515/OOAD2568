package model

import (
	commonModel "ModEd/common/model"

	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseUpdateRequest struct {
	gorm.Model
	CourseId    uuid.UUID                   `gorm:"type:varchar(255);not null;index"`
	RequestType CourseUpdateRequestTypeEnum `gorm:"type:string;not null"`
	UpdatedBy   commonModel.Instructor      `gorm:"foreignKey:UpdatedByID"`
	UpdatedAt   time.Time                   `gorm:"not null;autoUpdateTime"`
}

type CourseNameUpdate struct {
	gorm.Model
	CourseId  uuid.UUID              `gorm:"type:varchar(255);not null;index"`
	NewName   string                 `gorm:"type:string;not null"`
	UpdatedBy commonModel.Instructor `gorm:"foreignKey:UpdatedByID"`
}

type CoursePrerequisiteUpdate struct {
	gorm.Model
	CourseId  uuid.UUID              `gorm:"type:varchar(255);not null;index"`
	UpdatedBy commonModel.Instructor `gorm:"foreignKey:UpdatedByID"`
}
