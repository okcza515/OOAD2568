package model

import (
	"gorm.io/gorm"
)

type InternshipReport struct {
	gorm.Model
	InternshipReportID int    `gorm:"primaryKey	autoIncrement"`
	InternStudentID    int    `gorm:"foreignKey:InternID;references:InternStudentID"`
	ReportDate         string `gorm:"type:date"`
	ReportContent      string `gorm:"type:text"`
	ReportScore        int    `gorm:"type:int"`
}
