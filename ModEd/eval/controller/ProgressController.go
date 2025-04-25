package controller

import (
	commonModel "ModEd/common/model"

	evalModel "ModEd/eval/model"

	"time"

	"gorm.io/gorm"
)

type Progress struct {
	gorm.Model
	StudentCode  commonModel.Student `gorm:"foreignKey:StudentCode;references:StudentCode"`
	Title        evalModel.Assignment
	AssignmentId evalModel.Assignment
	Status       evalModel.Assignment
	LastUpdate   time.Time `gorm:"autoUpdateTime"`
	TotalSubmit  uint
}

type ProgressController struct {
	db *gorm.DB
}

func NewProgressController(db *gorm.DB) *ProgressController {
	return &ProgressController{db: db}
}

func (controller *ProgressController) GetAllProgress(AssignmentId uint) ([]Progress, error) {
	var progressList []Progress

	if err := controller.db.Where("assignment_id = ?", AssignmentId).
		Preload("Student").Preload("Title").
		Find(&progressList).Error; err != nil {
		return nil, err
	}
	return progressList, nil
}

func (controller *ProgressController) GetProgressByStudentCode(AssignmentId uint, StudentCode string) ([]Progress, error) {
	var progressList []Progress

	if err := controller.db.Where("assignment_id = ? AND student_code = ?", AssignmentId, StudentCode).
		Preload("Student").Preload("Title").
		Find(&progressList).Error; err != nil {
		return nil, err
	}
	return progressList, nil
}

func (controller *ProgressController) GetProgressByStatus(AssignmentId uint, status string) ([]Progress, error) {
	var progressList []Progress

	if err := controller.db.Where("assignment_id = ? AND status = ?", AssignmentId, status).
		Preload("Student").Preload("Title").
		Find(&progressList).Error; err != nil {
		return nil, err
	}
	return progressList, nil
}

func (controller *ProgressController) GetSubmitCountByAssignmentID(assignmentId uint) (uint, error) {
	var count int64

	if err := controller.db.Model(&evalModel.AssignmentSubmission{}).
		Where("assignment_id = ? AND submitted = ?", assignmentId, true).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return uint(count), nil
}
