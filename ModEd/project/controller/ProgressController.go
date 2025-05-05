package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"

	"gorm.io/gorm"
)

type ProgressController struct {
	*core.BaseController[*model.Progress]
	db *gorm.DB
}

func NewProgressController(db *gorm.DB) *ProgressController {
	return &ProgressController{
		db:             db,
		BaseController: core.NewBaseController[*model.Progress](db),
	}
}

func (c *ProgressController) ListAllProgress() ([]model.Progress, error) {
	progressPtrs, err := c.List(nil)
	if err != nil {
		return nil, err
	}

	progresses := make([]model.Progress, len(progressPtrs))
	for i, pPtr := range progressPtrs {
		progresses[i] = *pPtr
	}

	return progresses, nil
}

func (c *ProgressController) AddNewProgress(assignmentID uint, name string) error {
	progress := model.Progress{
		AssignmentId: assignmentID,
		Name:         name,
		IsCompleted:  false,
	}

	return c.Insert(&progress)
}

func (c *ProgressController) UpdateProgressName(progressID uint, newName string) error {
	progress, err := c.RetrieveByID(progressID)
	if err != nil {
		return fmt.Errorf("progress not found: %w", err)
	}

	progress.Name = newName
	return c.UpdateByID(progress)
}

func (c *ProgressController) GetFormattedProgressList() ([]string, error) {
	progressList, err := c.ListAllProgress()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve progress list: %w", err)
	}

	var formattedList []string
	for _, progress := range progressList {
		formattedList = append(formattedList, fmt.Sprintf(
			"Progress ID: %d, Assignment ID: %d, Name: %s, Completed: %t",
			progress.ID, progress.AssignmentId, progress.Name, progress.IsCompleted,
		))
	}

	return formattedList, nil
}

func (c *ProgressController) MarkAsCompleted(id uint) error {
	progress, err := c.RetrieveByID(id)
	if err != nil {
		return fmt.Errorf("progress not found: %w", err)
	}

	progress.IsCompleted = true
	return c.UpdateByID(progress)
}

func (c *ProgressController) MarkProgressAsIncomplete(progressID uint) error {
	progress, err := c.RetrieveByID(progressID)
	if err != nil {
		return fmt.Errorf("progress not found: %w", err)
	}

	progress.IsCompleted = false
	return c.UpdateByID(progress)
}
