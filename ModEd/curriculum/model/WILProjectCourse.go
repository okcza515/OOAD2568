// MEP-1010 Work Integrated Learning (WIL)
package model

type WILProjectCourse struct {
	CourseId uint   `gorm:"not null" validation:"required,uint"`
	Semester string `gorm:"not null" validation:"required,uint"`
}
