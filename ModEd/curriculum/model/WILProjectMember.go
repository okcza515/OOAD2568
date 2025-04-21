package model

import (
	commonModel "ModEd/common/model"
)

type WILProjectMember struct {
	WILProjectApplicationId uint                `json:"-" gorm:"primaryKey"`
	StudentId               string              `json:"StudentId" gorm:"primaryKey"`
	Student                 commonModel.Student `json:"Student" gorm:"foreignKey:StudentId;references:StudentCode"`
}
