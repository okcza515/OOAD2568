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

type Curriculum struct {
	CurriculumId uint              `gorm:"primaryKey;unique" csv:"curriculum_id" json:"curriculum_id" validate:"required"`
	Name         string            `gorm:"not null" csv:"name" json:"name" validate:"required"`
	StartYear    int               `gorm:"not null" csv:"start_year" json:"start_year" validate:"required"`
	EndYear      int               `gorm:"not null" csv:"end_year" json:"end_year" validate:"required,gtefield=StartYear"`
	DepartmentId uint              `gorm:"not null" csv:"department_id" json:"department_id" validate:"required"`
	Department   model.Department  `gorm:"foreignKey:ID;references:DepartmentId" csv:"-" json:"-" validate:"-"`
	ProgramType  model.ProgramType `gorm:"type:text;not null" csv:"program_type" json:"program_type" validate:"required"`
	CourseList   []Course          `gorm:"foreignKey:CurriculumId;references:CurriculumId" csv:"-" json:"-" validate:"-"`
	CreatedAt    time.Time         `gorm:"autoCreateTime" csv:"created_at" json:"created_at" validate:"-"`
	UpdatedAt    time.Time         `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at" validate:"-"`
	DeletedAt    gorm.DeletedAt    `csv:"-" json:"-" validate:"-"`
	*core.SerializableRecord
}

func (c *Curriculum) GetID() uint {
	return c.CurriculumId
}
func (c *Curriculum) ToString() string {
	return fmt.Sprintf("%+v", c)
}
func (c *Curriculum) Validate() error {
	validate := validator.New()

	// Validate struct fields using v10 validator
	if err := validate.Struct(c); err != nil {
		return err
	}
	return nil
}

// Testing functions
func (c *Curriculum) Print() {
	fmt.Println("───────────────────────────────────────────────────────────────────────────")
	fmt.Printf("CURRICULUM DETAILS %-58s\n", "")
	fmt.Println("───────────────────────────────────────────────────────────────────────────")
	fmt.Printf("ID:          %-60d\n", c.CurriculumId)
	fmt.Printf("Name:        %-60s\n", c.Name)
	fmt.Printf("Start Year:  %-60d\n", c.StartYear)
	fmt.Printf("End Year:    %-60d\n", c.EndYear)
	fmt.Printf("Department Id: %-60d\n", c.DepartmentId)
	fmt.Printf("Program Type: %-60s\n", c.ProgramType)
	fmt.Printf("Created At:  %-60s\n", c.CreatedAt)
	fmt.Printf("Updated At:  %-60s\n", c.UpdatedAt)
	fmt.Println("───────────────────────────────────────────────────────────────────────────")
}
