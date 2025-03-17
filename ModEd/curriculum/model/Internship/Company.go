package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	CompanyId   uuid.UUID
	CompanyName string
}
