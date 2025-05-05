// MEP-1010 Work Integrated Learning (WIL)
package model

import (
	"ModEd/core"
)

type WILProject struct {
	core.BaseModel
	ClassId         uint   `gorm:"not null" validate:"required"`
	SeniorProjectId uint   `gorm:"not null" validate:"required"`
	Company         uint   `gorm:"not null" validate:"required"`
	Mentor          string `gorm:"not null" validate:"required"`
}
