package model

import (
	"gorm.io/gorm"
)

type InternshipReport struct {
	gorm.Model
	ReportScore   int    `gorm:"type:int"`
}
