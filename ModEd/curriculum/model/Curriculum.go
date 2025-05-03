// MEP-1002
package model

import (
	"ModEd/common/model"
	"ModEd/core"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Curriculum struct {
	CurriculumId uint   `gorm:"primaryKey;unique" csv:"curriculum_id" json:"curriculum_id"`
	Name         string `gorm:"not null" csv:"name" json:"name"`
	StartYear    int    `gorm:"not null" csv:"start_year" json:"start_year"`
	EndYear      int    `gorm:"not null" csv:"end_year" json:"end_year"`

	DepartmentId uint              `gorm:"not null" csv:"department_id" json:"department_id"`
	Department   model.Department  `gorm:"foreignKey:ID;references:DepartmentId" csv:"-" json:"-"`
	ProgramType  model.ProgramType `gorm:"type:text;not null" csv:"program_type" json:"program_type"`
	CourseList   []Course          `gorm:"foreignKey:CurriculumId;references:CurriculumId" csv:"-" json:"-"`
	CreatedAt    time.Time         `gorm:"autoCreateTime" csv:"created_at" json:"created_at"`
	UpdatedAt    time.Time         `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt    `csv:"-" json:"-"`
	*core.SerializableRecord
}

func (c *Curriculum) GetID() uint {
	return c.CurriculumId
}
func (c *Curriculum) ToString() string {
	return fmt.Sprintf("%+v", c)
}
func (c *Curriculum) Validate() error {
	if c.CurriculumId == 0 {
		return fmt.Errorf("Curriculum ID cannot be zero")
	}
	if c.Name == "" {
		return fmt.Errorf("Curriculum Name cannot be empty")
	}

	if c.StartYear > c.EndYear {
		return fmt.Errorf("Start Year cannot be greater than End Year")
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
