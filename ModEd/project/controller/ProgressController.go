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

func (c *ProgressController) RetrieveProgress(id uint) (*model.Progress, error) {
	return c.RetrieveByID(id)
}

func (c *ProgressController) InsertProgress(progress model.Progress) error {
	progressCopy := progress
	return c.Insert(&progressCopy)
}

func (c *ProgressController) UpdateProgress(progress *model.Progress) error {
	return c.UpdateByID(progress)
}

func (c *ProgressController) DeleteProgress(id uint) error {
	return c.DeleteByID(id)
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
	progress, err := c.RetrieveProgress(id)
	if err != nil {
		return fmt.Errorf("progress not found: %w", err)
	}

	progress.IsCompleted = true
	return c.UpdateProgress(progress)
}
