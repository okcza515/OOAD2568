// MEP-1002
package model

import (
	"ModEd/core"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Course struct {
	CourseId     uint           `gorm:"primaryKey" csv:"course_id" json:"course_id"`
	Name         string         `gorm:"not null" csv:"name" json:"name"`
	Description  string         `gorm:"not null" csv:"description" json:"description"`
	CurriculumId uint           `gorm:"not null" csv:"curriculum_id" json:"curriculum_id"`
	Curriculum   Curriculum     `gorm:"foreignKey:CurriculumId;references:CurriculumId" csv:"-" json:"-"`
	Optional     bool           `gorm:"not null" csv:"optional" json:"optional"`
	CourseStatus CourseStatus   `gorm:"type:int;not null" csv:"course_status" json:"course_status"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" csv:"created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `csv:"-" json:"-"`
	ClassList    []Class        `gorm:"foreignKey:CourseId;references:CourseId" csv:"-" json:"-"`
	Prerequisite []Course       `gorm:"many2many:course_prerequisites;foreignKey:CourseId;joinForeignKey:CourseId;References:CourseId;joinReferences:PrerequisiteId" csv:"-" json:"-"`
	*core.SerializableRecord
}

func (c *Course) GetID() uint {
	return c.CourseId
}
func (c *Course) ToString() string {
	return fmt.Sprintf("%+v", c)
}
func (c *Course) Validate() error {
	if c.CourseId == 0 {
		return fmt.Errorf("Course ID cannot be zero")
	}
	if c.Name == "" {
		return fmt.Errorf("Course Name cannot be empty")
	}
	if c.Description == "" {
		return fmt.Errorf("Course Description cannot be empty")
	}
	if c.CurriculumId == 0 {
		return fmt.Errorf("Curriculum ID cannot be zero")
	}
	return nil
}

// Testing functions
func (c *Course) Print() {
	fmt.Printf("Course Id: %d, Name: %s, Description: %s, Curriculum Id: %d, Optional: %t, Course Status: %d\n",
		c.CourseId, c.Name, c.Description, c.CurriculumId, c.Optional, c.CourseStatus)
}
