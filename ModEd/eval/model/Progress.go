// Sawitt Ngamvilaisiriwong 65070503469
// MEP-1006

package model

import (
	"ModEd/common/model"

	"gorm.io/gorm"

	"time"
)

type Progress struct {
	gorm.Model
	SID         model.Student `gorm:"foreignKey:SID"`
	Title       Assignment
	Status      string
	LastUpdate  time.Time `gorm:"autoUpdateTime"`
	TotalSubmit uint
}
