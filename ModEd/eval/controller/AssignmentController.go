package controller

import (
	"ModEd/project/model"

	"gorm.io/gorm"
)

type IAssignmentController interface {
	ListAllAssignments() ([]model.Assignment, error)
	RetrieveAssignment(id uint) (*model.Assignment, error)
	InsertAssignment(assignment *model.Assignment) error
	UpdateAssignment(assignment *model.Assignment) error
	DeleteAssignment(id uint) error
}

type AssignmentController struct {
	db *gorm.DB
}

func NewAssignmentController(db *gorm.DB) IAssignmentController {
	return &AssignmentController{db: db}
}

func (c *AssignmentController) ListAllAssignments() ([]model.Assignment, error) {
	var assignments []model.Assignment
	err := c.db.Preload("Submission").Find(&assignments).Error
	return assignments, err
}

func (c *AssignmentController) RetrieveAssignment(id uint) (*model.Assignment, error) {
	var assignment model.Assignment
	err := c.db.Preload("Submission").First(&assignment, id).Error
	if err != nil {
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
	return c.db.Delete(&model.Assignment{}, id).Error
}
