package model

type WILProjectClass struct {
	CourseId uint `gorm:"not null"`
	ClassId  uint `gorm:"not null"`
}
