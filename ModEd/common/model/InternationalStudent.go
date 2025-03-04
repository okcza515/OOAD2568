package model

import "gorm.io/gorm"

type InternationalStudent struct {
	gorm.Model
	Student
}
