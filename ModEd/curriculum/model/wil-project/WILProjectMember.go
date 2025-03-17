package model

import (
	"ModEd/common/model"

	"github.com/google/uuid"
)

type WILProjectMember struct {
	WILProjectApplicationId uuid.UUID     `json:"-" gorm:"primaryKey"`
	StudentId               string        `json:"StudentId" gorm:"primaryKey"`
	Student                 model.Student `json:"Student" gorm:"foreignKey:StudentId;references:SID"`
}
