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
	StudentCode model.Student `gorm:"foreignKey:StudentCode;references:StudentCode"`
	Title       Assignment
	Status      string
	LastUpdate  time.Time `gorm:"autoUpdateTime"`
	TotalSubmit uint
}
