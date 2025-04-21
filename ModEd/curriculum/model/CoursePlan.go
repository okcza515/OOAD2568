// MEP-1008
package model

import (
	common "ModEd/common/model"

	"time"

	"gorm.io/gorm"
)

type CoursePlan struct {
	gorm.Model

	CourseId uint   `gorm:"not null" json:"course_id"`
	Course   Course `gorm:"foreignKey:CourseId;references:CourseId" json:"-"`

	Week uint      `gorm:"not null" csv:"week" json:"week"`
	Date time.Time `gorm:"not null" json:"date"`

	InstructorId uint              `gorm:"not null" json:"instructor_id"`
	Instructor   common.Instructor `gorm:"type:varchar(100);not null" json:"instructor"`
	Topic        string            `gorm:"type:varchar(255);not null" json:"topic"`
	Description  string            `gorm:"type:text" json:"description"`
}
