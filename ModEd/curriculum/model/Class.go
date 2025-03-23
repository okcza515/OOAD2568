package model

import (
	"ModEd/common/model"
	"time"

	"gorm.io/gorm"
)

type Class struct {
	ClassId     uint            `gorm:"primaryKey" csv:"class_id" json:"class_id"`
	CourseId    uint            `gorm:"not null" csv:"course_id" json:"course_id"`
	Course      Course          `gorm:"foreignKey:CourseId;references:CourseId" csv:"-" json:"-"`
	Section     int             `gorm:"not null" csv:"section" json:"section"`
	Schedule    time.Time       `gorm:"not null" csv:"schedule" json:"schedule"`
	StudentList []model.Student `gorm:"many2many:class_students" csv:"-" json:"-"`
	// TODO: might need to create another field of InstructorIds for loading data from csv, json
	Instructors []model.Instructor `gorm:"many2many:class_instructors;" csv:"-" json:"-"`
	CreatedAt   time.Time          `gorm:"autoCreateTime" csv:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt     `csv:"-" json:"-"`
}
