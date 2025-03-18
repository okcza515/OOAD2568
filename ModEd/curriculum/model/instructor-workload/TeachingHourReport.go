package model

import (
	commonModel "ModEd/common/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TeacherHourReport struct {
	gorm.Model
	ReportId     uuid.UUID              `gorm:"type:string;not null;index"`
	instructorId commonModel.Instructor `gorm:"foreignKey:InstructorID"`
	TeachingHour float64                `gorm:"type:decimal(5,2);not null"`
}
