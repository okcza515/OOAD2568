package controller

import (
	"ModEd/project/model"

	"gorm.io/gorm"
)

type IGroupMemberController interface {
	ListAllGroupMembers() ([]model.GroupMember, error)
	RetrieveGroupMember(id uint) (*model.GroupMember, error)
	InsertGroupMember(GroupMember *model.GroupMember) error
	UpdateGroupMember(GroupMember *model.GroupMember) error
	DeleteGroupMember(id uint) error
}

type GroupMemberController struct {
	db *gorm.DB
}

func NewGroupMemberController(db *gorm.DB) IGroupMemberController {
	return &GroupMemberController{db: db}
}

func (c *GroupMemberController) ListAllGroupMembers() ([]model.GroupMember, error) {
	var GroupMembers []model.GroupMember
	err := c.db.Find(&GroupMembers).Error
	return GroupMembers, err
}

func (c *GroupMemberController) RetrieveGroupMember(id uint) (*model.GroupMember, error) {
	var GroupMember model.GroupMember
	if err := c.db.Where("id = ?", id).First(&GroupMember).Error; err != nil {
		return nil, err
	}
	return &GroupMember, nil
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
