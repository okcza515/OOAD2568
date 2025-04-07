package controller

import (
	"ModEd/core"
	"ModEd/project/model"
	"fmt"

	"gorm.io/gorm"
)

type ProgressController struct {
	*core.BaseController
	db *gorm.DB
}

func NewProgressController(db *gorm.DB) *ProgressController {
	return &ProgressController{
		db:             db,
		BaseController: core.NewBaseController("progresses", db),
	}
}

func (c *ProgressController) ListAllProgress() ([]core.RecordInterface, error) {
	return c.List(map[string]interface{}{})
}

func (c *ProgressController) RetrieveProgress(id uint) (*core.RecordInterface, error) {
	return c.RetrieveByID(id)
}

func (c *ProgressController) InsertProgress(progress model.Progress) error {
	return c.Insert(&progress)
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
	for _, record := range progressList {
		progress := record.(*model.Progress)
		formattedList = append(formattedList, fmt.Sprintf(
			"Progress ID: %d, Assignment ID: %d, Name: %s, Completed: %t",
			progress.ID, progress.AssignmentId, progress.Name, progress.IsCompleted,
		))
	}

	return formattedList, nil
}

func (c *ProgressController) MarkAsCompleted(id uint) error {
	var progress model.Progress
	if err := c.db.First(&progress, id).Error; err != nil {
		return fmt.Errorf("progress not found: %w", err)
	}

	progress.IsCompleted = true

	if err := c.db.Save(&progress).Error; err != nil {
		return fmt.Errorf("failed to mark progress as completed: %w", err)
	}
	return nil
}
