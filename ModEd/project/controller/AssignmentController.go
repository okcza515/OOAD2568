package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"
	"time"

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

func (c *AssignmentController) RetrieveAssignment(id uint) (*model.Assignment, error) {
	return c.RetrieveByID(id)
}

func (c *AssignmentController) RetrieveAssignmentsBySeniorProjectId(seniorProjectId uint) ([]*model.Assignment, error) {
	assignments, err := c.List(map[string]interface{}{"senior_project_id": seniorProjectId})
	if err != nil {
		return nil, err
	}
	return assignments, nil
}

func (c *AssignmentController) InsertAssignment(seniorProjectId uint, name, description string, dueDate time.Time) (*model.Assignment, error) {
	existing, err := c.List(map[string]interface{}{"senior_project_id": seniorProjectId})
	if err != nil {
		return nil, fmt.Errorf("failed to check existing assignments: %w", err)
	}
	if len(existing) > 0 {
		return nil, fmt.Errorf("assignment already exists for project %d", seniorProjectId)
	}
	if dueDate.IsZero() {
		return nil, fmt.Errorf("due date cannot be empty")
	}

	assignment := model.Assignment{
		Name:            name,
		SeniorProjectId: seniorProjectId,
		Description:     description,
		DueDate:         dueDate,
	}

	return &assignment, c.Insert(&assignment)
}

func (c *AssignmentController) UpdateAssignment(assignment *model.Assignment) error {
	return c.UpdateByID(assignment)
}

func (c *AssignmentController) DeleteAssignment(id uint) error {
	return c.DeleteByID(id)
}
