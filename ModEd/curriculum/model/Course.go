// MEP-1002
package model

import (
	"ModEd/core"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Course struct {
	CourseId     uint           `gorm:"primaryKey" csv:"course_id" json:"course_id" validate:"required"`
	Name         string         `gorm:"not null" csv:"name" json:"name" validate:"required"`
	Description  string         `gorm:"not null" csv:"description" json:"description" validate:"required"`
	CurriculumId uint           `gorm:"not null" csv:"curriculum_id" json:"curriculum_id" validate:"required"`
	Curriculum   Curriculum     `gorm:"foreignKey:CurriculumId;references:CurriculumId" csv:"-" json:"-" validate:"-"`
	Optional     bool           `gorm:"not null" csv:"optional" json:"optional" validate:"required"`
	CourseStatus CourseStatus   `gorm:"type:int;not null" csv:"course_status" json:"course_status" validate:"required"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" csv:"created_at" json:"created_at" validate:"-"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at" validate:"-"`
	DeletedAt    gorm.DeletedAt `csv:"-" json:"-" validate:"-"`
	ClassList    []Class        `gorm:"foreignKey:CourseId;references:CourseId" csv:"-" json:"-" validate:"-"`
	Prerequisite []Course       `gorm:"many2many:course_prerequisites;foreignKey:CourseId;joinForeignKey:CourseId;References:CourseId;joinReferences:PrerequisiteId" csv:"-" json:"-" validate:"-"`
	*core.SerializableRecord
}

func (c *Course) GetID() uint {
	return c.CourseId
}
func (c *Course) ToString() string {
	return fmt.Sprintf("%+v", c)
}
func (c *Course) Validate() error {
	validate := validator.New()

	// Validate struct fields using v10 validator
	if err := validate.Struct(c); err != nil {
		return err
	}

	return nil
}

// Testing functions
func (c *Course) Print() {
	fmt.Println("───────────────────────────────────────────────────────────────────────────")
	fmt.Printf("COURSE DETAILS %-58s\n", "")
	fmt.Println("───────────────────────────────────────────────────────────────────────────")
	fmt.Printf("ID:          %-60d\n", c.CourseId)
	fmt.Printf("Name:        %-60s\n", c.Name)
	fmt.Printf("Description: %-60s\n", c.Description)
	fmt.Printf("Curriculum:  %-60d\n", c.CurriculumId)
	fmt.Printf("Optional:    %-60t\n", c.Optional)
	fmt.Printf("Status:      %-60s\n", CourseStatusLabel[c.CourseStatus])
	fmt.Printf("Created At:  %-60s\n", c.CreatedAt)
	fmt.Printf("Updated At:  %-60s\n", c.UpdatedAt)
	fmt.Println("───────────────────────────────────────────────────────────────────────────")
}
