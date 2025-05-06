package controller

import (
	"ModEd/core"
	"ModEd/project/model"

	"gorm.io/gorm"
)

type GroupMemberController struct {
	*core.BaseController[*model.GroupMember]
	DB *gorm.DB
}

func NewGroupMemberController(db *gorm.DB) *GroupMemberController {
	return &GroupMemberController{
		BaseController: core.NewBaseController[*model.GroupMember](db),
		DB:             db,
	}
}

func (c *GroupMemberController) RetrieveGroupMembersBySeniorProjectId(projectId uint) ([]*model.GroupMember, error) {
	return c.List(map[string]interface{}{"senior_project_id": projectId})
}
