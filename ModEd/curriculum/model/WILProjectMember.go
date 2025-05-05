// MEP-1010 Work Integrated Learning (WIL)
package model

import (
	commonModel "ModEd/common/model"
	"ModEd/core"
)

type WILProjectMember struct {
	core.BaseModel
	WILProjectApplicationId uint                `json:"-" gorm:"uniqueIndex:idx_wil_project_member" validation:"required,uint"`
	StudentId               string              `json:"StudentId" gorm:"uniqueIndex:idx_wil_project_member" validation:"required"`
	Student                 commonModel.Student `json:"Student" gorm:"foreignKey:StudentId;references:StudentCode"`
}
