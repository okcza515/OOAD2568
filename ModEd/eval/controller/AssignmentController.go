package controller

import (
	"errors"
	"time"

	assignmentModel "ModEd/eval/model"

	"gorm.io/gorm"
)

type AssignmentController interface {
	GetAll() ([]assignmentModel.Assignment, error)
	GetByID(id uint) (*assignmentModel.Assignment, error)
	Create(input AssignmentInput) (*assignmentModel.Assignment, error)
	Update(id uint, input AssignmentInput) (*assignmentModel.Assignment, error)
	Delete(id uint) error
}

type assignmentController struct {
	db *gorm.DB
}

func NewAssignmentController(db *gorm.DB) AssignmentController {
	return &assignmentController{db: db}
}

type AssignmentInput struct {
	Title       string
	Description string
	StartDate   time.Time
	DueDate     time.Time
	Status      string
}

func (c *assignmentController) GetAll() ([]assignmentModel.Assignment, error) {
	var assignments []assignmentModel.Assignment
	if err := c.db.Preload("Submission").Find(&assignments).Error; err != nil {
		return nil, err
	}
	return assignments, nil
}

func (c *assignmentController) GetByID(id uint) (*assignmentModel.Assignment, error) {
	var assignment assignmentModel.Assignment
	if err := c.db.Preload("Submission").First(&assignment, id).Error; err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (c *assignmentController) Create(input AssignmentInput) (*assignmentModel.Assignment, error) {
	assignment := assignmentModel.Assignment{
		Title:       input.Title,
		Description: input.Description,
		StartDate:   input.StartDate,
		DueDate:     input.DueDate,
		Status:      input.Status,
	}
	if err := c.db.Create(&assignment).Error; err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (c *assignmentController) Update(id uint, input AssignmentInput) (*assignmentModel.Assignment, error) {
	var assignment assignmentModel.Assignment
	if err := c.db.First(&assignment, id).Error; err != nil {
		return nil, err
	}

	assignment.Title = input.Title
	assignment.Description = input.Description
	assignment.StartDate = input.StartDate
	assignment.DueDate = input.DueDate
	assignment.Status = input.Status

	if err := c.db.Save(&assignment).Error; err != nil {
		return nil, err
	}
	return &assignment, nil
}

func (c *assignmentController) Delete(id uint) error {
	var assignment assignmentModel.Assignment
	if err := c.db.First(&assignment, id).Error; err != nil {
		return errors.New("assignment not found")
	}
	return c.db.Delete(&assignment).Error
}
