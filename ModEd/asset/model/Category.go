package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryID   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"category_id" csv:"category_id"`
	CategoryName string    `gorm:"type:varchar(255);not null"`
	Description  string    `gorm:"type:varchar(255)"`

	UpdatedAt time.Time      `gorm:"autoCreateTime;not null" json:"timestamp" csv:"timestamp"`
	UpdatedBy *uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"updated_by" csv:"updated_by"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty" csv:"deleted_at,omitempty"`
	DeletedBy *uuid.UUID     `gorm:"type:uuid" json:"deleted_by,omitempty" csv:"deleted_by,omitempty"`
}
