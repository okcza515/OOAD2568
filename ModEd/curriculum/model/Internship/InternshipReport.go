package model

import (
	"gorm.io/gorm"
)

type InternshipReport struct {
	gorm.Model
	ReportDate    string `gorm:"type:date"`
	ReportContent string `gorm:"type:text"`
	ReportScore   int    `gorm:"type:int"`
}
