package model

import (
	commonModel "ModEd/common/model"

	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseUpdateRequest struct {
	gorm.Model
	CourseId       uuid.UUID                   `gorm:"type:varchar(255);not null;index"`
	RequestType    CourseUpdateRequestTypeEnum `gorm:"type:string;not null"`
	UpdatedBy      commonModel.Instructor      `gorm:"foreignKey:UpdatedByID"`
	UpdatedAt      time.Time                   `gorm:"not null;autoUpdateTime"`
	UpdateStrategy CourseUpdateStrategy        `gorm:"-"`
}

type CourseNameUpdate struct {
	Connector     *gorm.DB
	CourseRequest *CourseUpdateRequest
}

type PrerequisiteUpdate struct {
	Connector     *gorm.DB
	CourseRequest *CourseUpdateRequest
}
