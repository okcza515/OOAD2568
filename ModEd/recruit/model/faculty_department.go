// MEP-1003 Student Recruitment
package model

type Faculty struct {
	FacultyID uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null"`
}

type Department struct {
	DepartmentID uint   `gorm:"primaryKey"`
	Name         string `gorm:"unique;not null"`
	FacultyID    uint   `gorm:"type:TEXT;not null"`
}
