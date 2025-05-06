// MEP-1010 Work Integrated Learning (WIL)
package model

import (
	"ModEd/core"
)

type WILProject struct {
	core.BaseModel
	Mentor          string `gorm:"not null"`
	ClassId         uint   `gorm:"not null"`
	SeniorProjectId uint   `gorm:"not null"`
	Company         uint   `gorm:"not null"`
}
