// MEP-1003 Student Recruitment
package model

// Faculty represents a faculty in the university
type Faculty struct {
	FacultyID uint `gorm:"type:text;primaryKey;"`
	Name      string    `gorm:"unique;not null"`
}

// Department represents a department within a faculty
type Department struct {
	DepartmentID uint `gorm:"type:TEXT;primaryKey"`
	Name         string    `gorm:"unique;not null"`
	FacultyID    uint `gorm:"type:TEXT;not null"`
}
