package controller

import (
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type ProgressController struct {
	db *gorm.DB
}

func NewProgressController(db *gorm.DB) *ProgressController {
	return &ProgressController{db: db}
}

func (controller *ProgressController) GetAllProgress() ([]model.Progress, error) {
	var progressList []model.Progress

	if err := controller.db.Preload("Student").Preload("Title").Find(&progressList).Error; err != nil {
		return nil, err
	}

	return progressList, nil
}

func (controller *ProgressController) GetProgressByStudentCode(studentCode string) ([]model.Progress, error) {
	var progressList []model.Progress

	if err := controller.db.Where("student_code = ?", studentCode).Preload("Title").Preload("Student").Find(&progressList).Error; err != nil {
		return nil, err
	}

	return progressList, nil
}

func (controller *ProgressController) GetProgressByStatus(Status string) ([]model.Progress, error) {
	var progressList []model.Progress

	if err := controller.db.Where("status = ?", Status).Preload("Title").Preload("Student").Find(&progressList).Error; err != nil {
		return nil, err
	}

	return progressList, nil
}
