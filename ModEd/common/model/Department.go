package model

import (
	"gorm.io/gorm"
)

type Department struct {
	gorm.Model
	Name        	string      	`gorm:"not null" csv:"name" json:"name"`
	Faculty      	Faculty     	`gorm:"foreignKey:ParentId" json:"parent"`
	Students    	[]Student   	`gorm:"foreignKey:DepartmentId" json:"students"`
	Instructors 	*[]Instructor 	`gorm:"foreignKey:DepartmentId" json:"instructors"`
	//course
	Budget      	int         	`gorm:"default:0" csv:"budget" json:"budget"`
}