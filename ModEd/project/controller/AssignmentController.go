package controller

import (
	"ModEd/project/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IAssignmentController interface {
	ListAllAssignments() ([]model.Assignment, error)
	RetrieveAssignment(id uuid.UUID) (*model.Assignment, error)
	InsertAssignment(assignment *model.Assignment) error
	UpdateAssignment(assignment *model.Assignment) error
	DeleteAssignment(id uuid.UUID) error
}

type AssignmentController struct {
	db *gorm.DB
}

func NewAssignmentController(db *gorm.DB) IAssignmentController {
	return &AssignmentController{db: db}
}

func (c *AssignmentController) ListAllAssignments() ([]model.Assignment, error) {
	var assignments []model.Assignment
	err := c.db.Find(&assignments).Error
	return assignments, err
}

func (c *AssignmentController) RetrieveAssignment(id uuid.UUID) (*model.Assignment, error) {
	var assignment model.Assignment
	if err := c.db.Where("assignment_id = ?", id).First(&assignment).Error; err != nil {
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

func (c *AssignmentController) DeleteAssignment(id uuid.UUID) error {
	return c.db.Where("assignment_id = ?", id).Delete(&model.Assignment{}).Error
}
