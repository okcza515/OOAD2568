// MEP-1008
package model

import (
	commonModel "ModEd/common/model"
	"time"

	"gorm.io/gorm"
)

type ClassLecture struct {
	gorm.Model
	ClassId      uint                   `gorm:"not null" json:"class_id"`
	Class        Class                  `gorm:"foreignKey:ClassId;references:ClassId" json:"-"`
	LectureName  string                 `gorm:"not null" json:"lecture_name"`
	InstructorId uint                   `gorm:"not null" json:"instructor_id"`
	Instructor   commonModel.Instructor `gorm:"foreignKey:InstructorId;references:ID" json:"-"`
	StartTime    time.Time              `gorm:"not null" json:"start_time"`
	EndTime      time.Time              `gorm:"not null" json:"end_time"`
}
