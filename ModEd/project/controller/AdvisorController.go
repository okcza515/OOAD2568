package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	utils "ModEd/project/utils"
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
		ProjectID:    projectId,
		InstructorID: instructorId,
		IsPrimary:    isPrimary,
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
	advisor, err := getAdvisorByID(ac.DB, advisorId)
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

func (ac *AdvisorController) RemoveAdvisor(advisorId uint) error {
	advisor, err := getAdvisorByID(ac.DB, advisorId)
	if err != nil {
		return err
	}

	if err := ac.DB.Delete(advisor).Error; err != nil {
		return errors.New("failed to remove advisor")
	}

	return nil
}

func (ac *AdvisorController) ListAdvisorsByProject(projectId int) ([]model.Advisor, error) {
	var advisors []model.Advisor
	err := ac.DB.Where("seniorProjectId = ?", projectId).Find(&advisors).Error
	return advisors, err
}

func (ac *AdvisorController) ListAdvisorsByInstructor(instructorId int) ([]model.Advisor, error) {
	var advisors []model.Advisor
	err := ac.DB.Where("instructorId = ?", instructorId).Find(&advisors).Error
	return advisors, err
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

func getAdvisorByID(db *gorm.DB, advisorId uint) (*model.Advisor, error) {
	var advisor model.Advisor
	if err := db.First(&advisor, "advisorId = ?", advisorId).Error; err != nil {
		return nil, errors.New("advisor not found")
	}
	return &advisor, nil
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
