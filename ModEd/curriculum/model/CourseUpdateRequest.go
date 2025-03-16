package model

import (
	"time"

	"github.com/google/uuid"
)

type CourseUpdateRequest struct {
	CourseID    uuid.UUID                   `json:"course_id"`
	RequestType CourseUpdateRequestTypeEnum `json:"request_type"`
	// UpdatedBy   commonModel.Instructor `json:"updated_by"` wait for Instructor implemented
	UpdatedAt time.Time `json:"updated_at"`
}
