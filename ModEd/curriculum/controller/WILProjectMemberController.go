// MEP-1010 Work Integrated Learning (WIL)
package controller

import (
	"ModEd/core"
	"ModEd/curriculum/model"

	"gorm.io/gorm"
)

type WILProjectMemberController struct {
	connector *gorm.DB
	*core.BaseController[model.WILProjectMember]
}

type WILProjectMemberControllerInterface interface {
	InsertMany(data []model.WILProjectMember) error
}

func CreateWILProjectMemberController(connector *gorm.DB) *WILProjectMemberController {
	return &WILProjectMemberController{
		connector:      connector,
		BaseController: core.NewBaseController[model.WILProjectMember](connector),
	}
}
