package model

type WILProjectCourse struct {
	CourseId uint   `gorm:"not null"`
	Semester string `gorm:"not null"`
}
