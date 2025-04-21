package controller

import (
	"ModEd/core"
	"ModEd/project/model"

	"gorm.io/gorm"
)

type CommitteeController struct {
	*core.BaseController
	DB *gorm.DB
}

func NewCommitteeController(db *gorm.DB) *CommitteeController {
	return &CommitteeController{DB: db, BaseController: core.NewBaseController("committees", db)}
}

func (cc *CommitteeController) InsertCommittee(committee *model.Committee) error {
	err := cc.DB.Create(committee).Error
	if err != nil {
		return err
	}

	var project model.SeniorProject
	if err := cc.DB.First(&project, committee.SeniorProjectId).Error; err == nil {
	}
	return nil
}

func (cc *CommitteeController) ListCommitteesByProject(projectId int) ([]model.Committee, error) {
	var committees []model.Committee
	err := cc.DB.Where("seniorProjectId = ?", projectId).Find(&committees).Error
	return committees, err
}

func (cc *CommitteeController) ListCommitteesByInstructor(instructorId int) ([]model.Committee, error) {
	var committees []model.Committee
	err := cc.DB.Where("instructorId = ?", instructorId).Find(&committees).Error
	return committees, err
}

func (cc *CommitteeController) RemoveCommittee(committeeId int) error {
	return cc.DB.Delete(&model.Committee{}, committeeId).Error
}
