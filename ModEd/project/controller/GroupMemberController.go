package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"

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

func (c *GroupMemberController) ListAllGroupMembers() ([]*model.GroupMember, error) {
	groupMembers, err := c.List(map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	return groupMembers, nil
}

func (c *GroupMemberController) RetrieveGroupMember(id uint) (*model.GroupMember, error) {
	return c.RetrieveByID(id)
}

func (c *GroupMemberController) RetrieveGroupMembersBySeniorProjectId(projectId uint) ([]*model.GroupMember, error) {
	return c.List(map[string]interface{}{"senior_project_id": projectId})
}

func (c *GroupMemberController) InsertGroupMember(member *model.GroupMember) error {
	if member.SeniorProjectId == 0 {
		return fmt.Errorf("senior project ID is required")
	}
	return c.Insert(member)
}

func (c *GroupMemberController) UpdateGroupMember(member *model.GroupMember) error {
	return c.UpdateByID(member)
}

func (c *GroupMemberController) DeleteGroupMember(id uint) error {
	return c.DeleteByID(id)
}
