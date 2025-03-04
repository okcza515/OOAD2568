package model

import "gorm.io/gorm"

type RegularStudent struct {
	gorm.Model
	Student
}
