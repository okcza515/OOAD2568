package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	DepartmentId 	uuid.UUID   	`gorm:"type:uuid;primaryKey" csv:"department_id" json:"department_id"`
	Name        	string      	`gorm:"not null" csv:"name" json:"name"`
	Parent      	Faculty     	`gorm:"foreignKey:ParentId" json:"parent"`
	Students    	[]Student   	`gorm:"foreignKey:DepartmentId" json:"students"`
	Instructors 	[]Instructor 	`gorm:"foreignKey:DepartmentId" json:"instructors"`
	CourseId    	[]uuid.UUID 	`gorm:"type:uuid[]" csv:"course_id" json:"course_id"` // UUID to avoid circular dependency
	Budget      	int         	`gorm:"default:0" csv:"budget" json:"budget"`
}