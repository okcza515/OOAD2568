// MEP-1010 Work Integrated Learning (WIL)
package model

import (
	commonModel "ModEd/common/model"
	"ModEd/core"
)

type WILProjectMember struct {
	core.BaseModel
	WILProjectApplicationId uint                `json:"-" gorm:"uniqueIndex:idx_wil_project_member"`
	StudentId               string              `json:"StudentId" gorm:"uniqueIndex:idx_wil_project_member,not null" validation:"studentId"`
	Student                 commonModel.Student `json:"Student" gorm:"foreignKey:StudentId;references:StudentCode"`
}
