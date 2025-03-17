package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseUpdateRequest struct {
	CourseID    uuid.UUID
	RequestType CourseUpdateRequestTypeEnum
	UpdatedBy   Instructor
	UpdatedAt   time.Time

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
