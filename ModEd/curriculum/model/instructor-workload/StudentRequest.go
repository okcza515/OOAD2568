package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentRequest struct {
	gorm.Model
	RequestId uuid.UUID `gorm:"type:text;primaryKey;default:gen_random_uuid()"`
	StudentId uuid.UUID `gorm:"type:text;not null"`
	Remark    string    `gorm:"type:text"`
	Status    string    `gorm:"type:text;not null"`
	CreatedBy uuid.UUID `gorm:"type:text;not null"`
	CreatedAt string    `gorm:"type:date;not null"`
}
