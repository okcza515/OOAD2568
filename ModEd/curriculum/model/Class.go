// MEP-1002
package model

import (
	"ModEd/common/model"
	"ModEd/core"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Class struct {
	ClassId     uint               `gorm:"primaryKey" csv:"class_id" json:"class_id"`
	CourseId    uint               `gorm:"not null" csv:"course_id" json:"course_id"`
	Course      Course             `gorm:"foreignKey:CourseId;references:CourseId" csv:"-" json:"-"`
	Section     int                `gorm:"not null" csv:"section" json:"section"`
	Schedule    time.Time          `gorm:"not null" csv:"schedule" json:"schedule"`
	StudentList []model.Student    `gorm:"many2many:class_students" csv:"-" json:"-"`
	Instructors []model.Instructor `gorm:"many2many:class_instructors;" csv:"-" json:"-"`
	CreatedAt   time.Time          `gorm:"autoCreateTime" csv:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt     `csv:"-" json:"-"`
	*core.SerializableRecord
}

func (c *Class) GetID() uint {
	return c.ClassId
}

func (c *Class) ToString() string {
	return fmt.Sprintf("%+v", c)
}

func (c *Class) Validate() error {
	if c.ClassId == 0 {
		return fmt.Errorf("Class ID cannot be zero")
	}
	if c.CourseId == 0 {
		return fmt.Errorf("Course ID cannot be zero")
	}
	if c.Section <= 0 {
		return fmt.Errorf("Section must be greater than zero")
	}
	if c.Schedule.IsZero() {
		return fmt.Errorf("Schedule cannot be zero")
	}
	return nil
}

// Testing functions
func (c *Class) Print() {
	fmt.Println("───────────────────────────────────────────────────────────────────────────")
	fmt.Printf("CLASS DETAILS %-58s\n", "")
	fmt.Println("───────────────────────────────────────────────────────────────────────────")
	fmt.Printf("ID:          %-60d\n", c.ClassId)
	fmt.Printf("Course Id:   %-60d\n", c.CourseId)
	fmt.Printf("Section:     %-60d\n", c.Section)
	fmt.Printf("Schedule:    %-60s\n", c.Schedule)
	fmt.Printf("Created At:  %-60s\n", c.CreatedAt)
	fmt.Printf("Updated At:  %-60s\n", c.UpdatedAt)
	fmt.Println("───────────────────────────────────────────────────────────────────────────")
}
