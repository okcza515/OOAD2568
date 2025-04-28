package controller

import (
	"ModEd/core"
	"ModEd/project/model"

	"gorm.io/gorm"
)

type GroupMemberController struct {
	*core.BaseController
	db *gorm.DB
}

func NewGroupMemberController(db *gorm.DB) *GroupMemberController {
	return &GroupMemberController{
		db:             db,
		BaseController: core.NewBaseController("groupMembers", db),
	}
}

func (c *GroupMemberController) ListAllGroupMembers() ([]model.GroupMember, error) {
	var groupMembers []model.GroupMember
	err := c.db.Find(&groupMembers).Error
	return groupMembers, err
}

func (c *GroupMemberController) RetrieveGroupMember(id uint) (*model.GroupMember, error) {
	var groupMember model.GroupMember
	if err := c.db.Where("id = ?", id).First(&groupMember).Error; err != nil {
		return nil, err
	}
	return &groupMember, nil
}

func (c *GroupMemberController) InsertGroupMember(GroupMember *model.GroupMember) error {
	return c.db.Create(GroupMember).Error
}

func (c *GroupMemberController) UpdateGroupMember(GroupMember *model.GroupMember) error {
	return c.db.Save(GroupMember).Error
}

func (c *GroupMemberController) DeleteGroupMember(id uint) error {
	return c.db.Where("id = ?", id).Delete(&model.GroupMember{}).Error
}
