// MEP-1008
package model

import (
	commonModel "ModEd/common/model"
	"ModEd/core"

	"fmt"
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
	*core.SerializableRecord
}

func (c *ClassLecture) GetID() uint {
	return c.ClassId
}

func (c *ClassLecture) ToString() string {
	return fmt.Sprintf("%+v", c)
}

func (c *ClassLecture) Validate() error {
	if c.ClassId == 0 {
		return fmt.Errorf("Class ID cannot be zero")
	}
	if c.LectureName == "" {
		return fmt.Errorf("Lecture name cannot be empty")
	}
	if c.InstructorId == 0 {
		return fmt.Errorf("Instructor ID cannot be zero")
	}
	if c.StartTime.IsZero() {
		return fmt.Errorf("Start time cannot be empty")
	}
	if c.EndTime.IsZero() {
		return fmt.Errorf("End time cannot be empty")
	}
	if c.StartTime.After(c.EndTime) {
		return fmt.Errorf("Start time cannot be after end time")
	}
	return nil
}
