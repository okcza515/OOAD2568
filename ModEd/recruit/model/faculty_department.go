package model

import (
	"github.com/google/uuid"
)

// Faculty represents a faculty in the university
type Faculty struct {
	FacultyID uuid.UUID `gorm:"type:text;primaryKey;"`
	Name      string    `gorm:"unique;not null"`
}

// Department represents a department within a faculty
type Department struct {
	DepartmentID uuid.UUID `gorm:"type:TEXT;primaryKey"`
	Name         string    `gorm:"unique;not null"`
	FacultyID    uuid.UUID `gorm:"type:TEXT;not null"`
}
