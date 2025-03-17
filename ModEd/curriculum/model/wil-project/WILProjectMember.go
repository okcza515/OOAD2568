package model

import (
	"ModEd/common/model"
)

type WILProjectMember struct {
	WILProjectApplicationId string        `json:"-" gorm:"primaryKey"`
	StudentId               string        `json:"StudentId" gorm:"primaryKey"`
	Student                 model.Student `json:"Student" gorm:"foreignKey:StudentId;references:SID"`
}
