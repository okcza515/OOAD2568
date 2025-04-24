// MEP-1002
package model

import (
	"ModEd/common/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Curriculum struct {
	CurriculumId uint   `gorm:"primaryKey;unique" csv:"curriculum_id" json:"curriculum_id"`
	Name         string `gorm:"not null" csv:"name" json:"name"`
	StartYear    int    `gorm:"not null" csv:"start_year" json:"start_year"`
	EndYear      int    `gorm:"not null" csv:"end_year" json:"end_year"`
	// TODO: Removable DepartmentName?
	DepartmentName string `gorm:"not null" csv:"department_name" json:"department_name"`

	Department  model.Department  `gorm:"-" csv:"-" json:"-"`
	ProgramType model.ProgramType `gorm:"type:text;not null" csv:"program_type" json:"program_type"`
	CourseList  []Course          `gorm:"foreignKey:CurriculumId;references:CurriculumId" csv:"-" json:"-"`
	CreatedAt   time.Time         `gorm:"autoCreateTime" csv:"created_at" json:"created_at"`
	UpdatedAt   time.Time         `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt    `csv:"-" json:"-"`
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
	fmt.Printf("Curriculum Id: %d, Name: %s, Start Year: %d, End Year: %d, Department Name: %s, Program Type: %s\n",
		c.CurriculumId, c.Name, c.StartYear, c.EndYear, c.DepartmentName, c.ProgramType)
}
