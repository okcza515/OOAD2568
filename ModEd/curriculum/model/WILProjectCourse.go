// MEP-1010 Work Integrated Learning (WIL)
package model

type WILProjectCourse struct {
	CourseId uint   `gorm:"not null"`
	Semester string `gorm:"not null"`
}
