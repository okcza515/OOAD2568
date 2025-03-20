package model

import (
	"gorm.io/gorm"
)

type StudentRequest struct {
	gorm.Model
	RequestId int                      `gorm:"primaryKey;autoIncrement"`
	StudentId int                      `gorm:"type:text;not null"`
	Remark    string                   `gorm:"type:text"`
	Status    StudentRequestStatusEnum `gorm:"type:text;not null"`
	CreatedBy int                      `gorm:"type:text;not null"`
	CreatedAt string                   `gorm:"type:date;not null"`
}
