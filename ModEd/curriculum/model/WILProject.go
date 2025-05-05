// MEP-1010 Work Integrated Learning (WIL)
package model

import (
	"ModEd/core"
)

type WILProject struct {
	core.BaseModel
	Mentor          string `gorm:"not null" validation:"required"`
	ClassId         uint   `gorm:"not null" validation:"required,uint"`
	SeniorProjectId uint   `gorm:"not null" validation:"required,uint"`
	Company         uint   `gorm:"not null" validation:"required,uint"`
}
