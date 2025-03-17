package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CourseUpdateRequest struct {
	CourseID    uuid.UUID                   `json:"course_id"`
	RequestType CourseUpdateRequestTypeEnum `json:"request_type"`
	UpdatedBy   Instructor                  `json:"updated_by"`
	UpdatedAt   time.Time                   `json:"updated_at"`

	UpdateStrategy CourseUpdateStrategy `json:"-"`
}

type CourseNameUpdate struct {
	Connector     *gorm.DB
	CourseRequest *CourseUpdateRequest
}

type PrerequisiteUpdate struct {
	Connector     *gorm.DB
	CourseRequest *CourseUpdateRequest
}
