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
	// TODO: in department, we have faculty already, so we don't need to link with faculty again
	// FacultyId    uint          `gorm:"not null" csv:"faculty_id" json:"faculty_id"`
	// Faculty      model.Faculty `gorm:"foreignKey:FacultyId;references:FacultyId"`
	// TODO: Why we have to link with DepartmentName, can we modify Department struct ???, let just ignore department with gorm for now because the reference is not correct
	// DepartmentId uint             `gorm:"not null" csv:"department_id" json:"department_id"`
	// Department   model.Department `gorm:"foreignKey:DepartmentId;references:DepartmentId"`
	DepartmentName string `gorm:"not null" csv:"department_name" json:"department_name"`
	// Department     model.Department `gorm:"foreignKey:DepartmentName;references:Name" csv:"-" json:"-"`
	Department model.Department `gorm:"-" csv:"-" json:"-"`

	ProgramType model.ProgramType `gorm:"type:text;not null" csv:"program_type" json:"program_type"`
	CourseList  []Course          `gorm:"foreignKey:CurriculumId;references:CurriculumId" csv:"-" json:"-"`
	CreatedAt   time.Time         `gorm:"autoCreateTime" csv:"created_at" json:"created_at"`
	UpdatedAt   time.Time         `gorm:"autoUpdateTime" csv:"updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt    `csv:"-" json:"-"`
}

func (c *Curriculum) Print() {
	fmt.Printf("Curriculum Id: %d, Name: %s, Start Year: %d, End Year: %d, Department Name: %s, Program Type: %s\n",
		c.CurriculumId, c.Name, c.StartYear, c.EndYear, c.DepartmentName, c.ProgramType)
}
