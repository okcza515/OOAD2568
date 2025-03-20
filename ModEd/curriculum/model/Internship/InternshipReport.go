package model

import (
	"gorm.io/gorm"
)

type InternshipReport struct {
	gorm.Model
	InternshipReportID int           `gorm:"primaryKey	autoIncrement"`
	Student            InternStudent `gorm:"foreignKey:InternStudentId"`
	ReportDate         string        `gorm:"type:date"`
	ReportContent      string        `gorm:"type:text"`
	ReportScore        int           `gorm:"type:int"`
}
