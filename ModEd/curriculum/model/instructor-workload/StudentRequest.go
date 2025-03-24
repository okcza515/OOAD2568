package model

import (
	commonModel "ModEd/common/model"

	"gorm.io/gorm"
)

type StudentRequest struct {
	gorm.Model
	StudentId   int                      `gorm:"not null;index"`
	RequestType StudentRequestTypeEnum   `gorm:"type:ENUM('SickLeave', 'PersonalLeave', 'ExtraCourseEnrollment');not null"`
	Remark      string                   `gorm:"type:text"`
	Status      StudentRequestStatusEnum `gorm:"type:ENUM('Pending', 'Approved', 'Rejected');not null"`
	Student     commonModel.Student      `gorm:"foreignKey:StudentId;references:ID"`
}
