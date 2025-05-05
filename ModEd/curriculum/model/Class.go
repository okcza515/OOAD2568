// MEP-1002
package model

import (
	"ModEd/common/model"
	"ModEd/core"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Class struct {
	ClassId     uint               `gorm:"primaryKey" csv:"class_id" json:"class_id" validate:"required"`
	CourseId    uint               `gorm:"not null" csv:"course_id" json:"course_id" validate:"required"`
	Course      Course             `gorm:"foreignKey:CourseId;references:CourseId" csv:"-" json:"-" validate:"-"`
	Section     int                `gorm:"not null" csv:"section" json:"section" validate:"required,gt=0"`
	Schedule    time.Time          `gorm:"not null" csv:"schedule" json:"schedule" validate:"required"`
	StudentList []model.Student    `gorm:"many2many:class_students" csv:"-" json:"-" validate:"-"`
	Instructors []model.Instructor `gorm:"many2many:class_instructors;" csv:"-" json:"-" validate:"-"`
	CreatedAt   time.Time          `gorm:"autoCreateTime" csv:"created_at" json:"created_at" validate:"-"`
	UpdatedAt   time.Time          `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at" validate:"-"`
	DeletedAt   gorm.DeletedAt     `csv:"-" json:"-" validate:"-"`
	*core.SerializableRecord
}

func (c *Class) GetID() uint {
	return c.ClassId
}

func (c *Class) ToString() string {
	return fmt.Sprintf("%+v", c)
}

func (c *Class) Validate() error {
	validate := validator.New()

	// Validate struct fields using v10 validator
	if err := validate.Struct(c); err != nil {
		return err
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
