package model

import (
	"gorm.io/gorm"
)

type RequestResignationStudent struct {
	gorm.Model
	StudentCode string `gorm:"type:text;default:'';not null"` // อ้างถึง StudentInfo.StudentCode
	Reason      string `gorm:"type:text"`                     // optional เหตุผลลาออก
	Status      string `gorm:"default:Pending"`               // Pending / Approved / Rejected
}

func (r RequestResignationStudent) GetID() string {
	return r.StudentCode
}

func (r RequestResignationStudent) GetReason() string {
	return r.Reason
}

func (r RequestResignationStudent) GetStatus() string {
	return r.Status
}
