package model

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type InternshipReport struct {
	gorm.Model
	InternshipReportID 		uuid.UUID `gorm:"primaryKey"`
	InternshipID 					uuid.UUID	`gorm:"foreignKey:InternID;references:InternID"`
	ReportDate 						string		`gorm:"type:date"`
	ReportContent 				string		`gorm:"type:text"`
	ReportScore 					int				`gorm:"type:int"`
}