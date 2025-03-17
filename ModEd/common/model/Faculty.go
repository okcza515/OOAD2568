package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Faculty struct {
	gorm.Model
	FacultyId 	uuid.UUID
	Name 		string
	Budget 		int
}
