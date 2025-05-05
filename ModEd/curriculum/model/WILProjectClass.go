// MEP-1010 Work Integrated Learning (WIL)
package model

type WILProjectClass struct {
	CourseId uint `gorm:"not null" validation:"required,uint"`
	ClassId  uint `gorm:"not null" validation:"required,uint"`
}
