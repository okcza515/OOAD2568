package model

import (
	"gorm.io/gorm"
)

type Faculty struct {
	gorm.Model
	Name 	string
	Budget 	int
}
