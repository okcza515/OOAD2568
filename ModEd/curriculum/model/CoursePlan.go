// MEP-1008
package model

import (
	common "ModEd/common/model"
	"ModEd/core"

	"fmt"
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
	Instructor   common.Instructor `gorm:"-" json:"instructor"`
	Topic        string            `gorm:"type:varchar(255);not null" json:"topic"`
	Description  string            `gorm:"type:text" json:"description"`
	*core.SerializableRecord
}

func (c *CoursePlan) GetID() uint {
	return c.CourseId
}

func (c *CoursePlan) ToString() string {
	return fmt.Sprintf("%+v", c)
}

func (c *CoursePlan) Validate() error {
	if c.CourseId == 0 {
		return fmt.Errorf("Course ID cannot be zero")
	}
	if c.Week == 0 {
		return fmt.Errorf("Week cannot be zero")
	}
	if c.Date.IsZero() {
		return fmt.Errorf("Date cannot be empty")
	}
	if c.InstructorId == 0 {
		return fmt.Errorf("Instructor ID cannot be zero")
	}
	if c.Topic == "" {
		return fmt.Errorf("Topic cannot be empty")
	}
	return nil
}
