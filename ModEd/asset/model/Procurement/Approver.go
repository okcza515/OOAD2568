package model

import (
	"gorm.io/gorm"
)

type Approver struct {
	ApproverID uint `gorm:":primaryKey"`
	Name       string
	Department string
	Email      string
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
