package model

import (
	"gorm.io/gorm"
)

type RequestResignationInstructor struct {
	gorm.Model
	Reason    string `gorm:"type:text"`       // optional เหตุผลลาออก
	Status    string `gorm:"default:Pending"` // Pending / Approved / Rejected
}
