package model

import (
	"ModEd/common/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ClassId     uuid.UUID       `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" csv:"class_id" json:"class_id"`
	Course      Course          `csv:"-" json:"-"`
	CourseId    uuid.UUID       `csv:"course_id" json:"course_id"`
	Section     int             `csv:"section" json:"section"`
	Schedule    time.Time       `csv:"schedule" json:"schedule"`
	StudentList []model.Student `csv:"-" json:"-"`
	// TODO: Instructor is not defined in the model package, wait for Aj Bo
	// Instructor   model.Instructor
}
