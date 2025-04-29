package controller

import (
	"ModEd/core"
	"ModEd/project/model"

	"gorm.io/gorm"
)

type AssignmentController struct {
	*core.BaseController[*model.Assignment]
	db *gorm.DB
}

func NewAssignmentController(db *gorm.DB) *AssignmentController {
	return &AssignmentController{
		db:             db,
		BaseController: core.NewBaseController[*model.Assignment](db),
	}
}

func (c *AssignmentController) ListAllAssignments() ([]model.Assignment, error) {
	var assignments []model.Assignment
	result := c.db.Find(&assignments)
	if result.Error != nil {
		return nil, result.Error
	}
	return assignments, nil
}

func (c *AssignmentController) RetrieveAssignment(seniorProjectId uint) (*model.Assignment, error) {
	var assignment model.Assignment
	if err := c.db.First(&assignment, "senior_project_id = ?", seniorProjectId).Error; err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (c *AssignmentController) RetrieveAssignmentBySeniorProjectId(seniorProjectId uint) (*model.Assignment, error) {
	var assignment model.Assignment
	if err := c.db.First(&assignment, "senior_project_id = ?", seniorProjectId).Error; err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (c *AssignmentController) InsertAssignment(assignment *model.Assignment) error {
	return c.db.Create(assignment).Error
}

func (c *AssignmentController) UpdateAssignment(assignment *model.Assignment) error {
	return c.db.Save(assignment).Error
}

func (c *AssignmentController) DeleteAssignment(id uint) error {
	return c.db.Where("id = ?", id).Delete(&model.Assignment{}).Error
}
