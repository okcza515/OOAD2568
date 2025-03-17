package model

import (
	"ModEd/common/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Curriculum struct {
	gorm.Model
	CurriculumId uuid.UUID `gorm:"type:uuid;primaryKey" csv:"curriculum_id" json:"curriculum_id"`
	Name         string
	StartYear    int
	EndYear      int
	FacultyId    uuid.UUID
	Faculty      model.Faculty `gorm:"foreignKey:FacultyId;references:FacultyId"`
	DepartmentId uuid.UUID
	Department   model.Department `gorm:"foreignKey:DepartmentId;references:DepartmentId"`
	ProgramType  model.ProgramType
	CourseList   []Course
	CreatedAt    time.Time
	UpdatedAt    time.Time `gorm:"autoUpdateTime;column:updated_at"`
}
