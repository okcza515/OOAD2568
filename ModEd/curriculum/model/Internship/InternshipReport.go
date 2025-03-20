package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InternshipReport struct {
	gorm.Model
	InternshipReportID int       `gorm:"primaryKey	autoIncrement"`
	InternshipID       uuid.UUID `gorm:"foreignKey:InternID;references:InternID"`
	ReportDate         string    `gorm:"type:date"`
	ReportContent      string    `gorm:"type:text"`
	ReportScore        int       `gorm:"type:int"`
}
