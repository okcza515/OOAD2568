// MEP-1008
package model

import (
	common "ModEd/common/model"
	"ModEd/core"
	"fmt"

	gorm "gorm.io/gorm"
)

type StudentAdvisor struct {
	gorm.Model
	InstructorId uint              `gorm:"not null" json:"instructor_id"`
	Instructor   common.Instructor `gorm:"foreignKey:InstructorId;references:ID" json:"-"`
	Students     []common.Student  `gorm:"many2many:student_advisor_students;"`
	*core.SerializableRecord
}

func (sa *StudentAdvisor) GetID() uint {
	return sa.ID
}

func (sa *StudentAdvisor) ToString() string {
	return fmt.Sprintf("%+v", sa)
}

func (sa *StudentAdvisor) Validate() error {
	if sa.InstructorId == 0 {
		return fmt.Errorf("Instructor ID cannot be zero")
	}
	if len(sa.Students) == 0 {
		return fmt.Errorf("Student list cannot be empty")
	}
	return nil
}
