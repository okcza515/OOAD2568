package model

import (
	"ModEd/common/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ClassId     uuid.UUID `gorm:"type:uuid;primaryKey" csv:"class_id" json:"class_id"`
	CourseId    uuid.UUID `csv:"course_id" json:"course_id"`
	Course      Course    `gorm:"foreignKey:CourseId;references:CourseId" csv:"-" json:"-"`
	Section     int       `csv:"section" json:"section"`
	Schedule    time.Time `csv:"schedule" json:"schedule"`
	StudentList []model.Student
	CreatedAt   time.Time
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	// TODO: Instructor is not defined in the model package, wait for Aj Bo
	// Instructor   model.Instructor
}
