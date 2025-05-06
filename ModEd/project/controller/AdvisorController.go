package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"errors"

	"gorm.io/gorm"
)

type AdvisorController struct {
	*core.BaseController[*model.Advisor]
	DB *gorm.DB
}

func NewAdvisorController(db *gorm.DB) *AdvisorController {
	return &AdvisorController{
		BaseController: core.NewBaseController[*model.Advisor](db),
		DB:             db,
	}
}

func (ac *AdvisorController) AssignAdvisor(projectId, instructorId uint, isPrimary bool) (*model.Advisor, error) {
	if err := validateAdvisorAssignment(ac.DB, projectId, instructorId, isPrimary); err != nil {
		return nil, err
	}

	advisor := &model.Advisor{
		SeniorProjectId: projectId,
		InstructorId:    instructorId,
		IsPrimary:       isPrimary,
	}

	if err := ac.DB.Create(advisor).Error; err != nil {
		return nil, errors.New("failed to assign advisor")
	}

	var project model.SeniorProject
	if err := ac.DB.First(&project, projectId).Error; err == nil {
	}

	return advisor, nil
}

func (ac *AdvisorController) UpdateAdvisorRole(advisorId uint, isPrimary bool) error {
	advisor, err := ac.RetrieveByID(advisorId)
	if err != nil {
		return err
	}

	if isPrimary {
		if err := validatePrimaryAdvisorChange(ac.DB, advisor.SeniorProjectId, advisorId); err != nil {
			return err
		}
	}

	return ac.DB.Model(advisor).Update("isPrimary", isPrimary).Error
}

func advisorExists(db *gorm.DB, projectId, instructorId uint) (bool, error) {
	var exists bool
	err := db.Raw(`
		SELECT EXISTS (
			SELECT 1 FROM advisors 
			WHERE seniorProjectId = ? AND instructorId = ?
		)
	`, projectId, instructorId).Scan(&exists).Error
	return exists, err
}

func hasPrimaryAdvisor(db *gorm.DB, projectId uint) (bool, error) {
	var exists bool
	err := db.Raw(`
		SELECT EXISTS (
			SELECT 1 FROM advisors 
			WHERE seniorProjectId = ? AND isPrimary = true
		)
	`, projectId).Scan(&exists).Error
	return exists, err
}

func validateAdvisorAssignment(db *gorm.DB, projectId, instructorId uint, isPrimary bool) error {
	if exists, _ := advisorExists(db, projectId, instructorId); exists {
		return errors.New("instructor is already assigned to this project")
	}

	if isPrimary {
		if hasPrimary, _ := hasPrimaryAdvisor(db, projectId); hasPrimary {
			return errors.New("a primary advisor is already assigned to this project")
		}
	}

	return nil
}

func validatePrimaryAdvisorChange(db *gorm.DB, projectId, advisorId uint) error {
	hasPrimary, err := hasPrimaryAdvisor(db, projectId)
	if err != nil {
		return err
	}
	if hasPrimary {
		return errors.New("a primary advisor is already assigned to this project")
	}
	return nil
}
