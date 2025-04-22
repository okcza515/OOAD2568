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

func (controller *ProgressController) GetAllProgress() ([]Progress, error) {
	var progressList []Progress

	if err := controller.db.Preload("Student").Preload("Title").Find(&progressList).Error; err != nil {
		return nil, err
	}
	return progressList, nil
}

func (controller *ProgressController) GetProgressByStudentCode(studentCode string) ([]Progress, error) {
	var progressList []Progress

	if err := controller.db.Where("student_code = ?", studentCode).
		Preload("Student").Preload("Title").
		Find(&progressList).Error; err != nil {
		return nil, err
	}
	return progressList, nil
}

func (controller *ProgressController) GetProgressByStatus(status string) ([]Progress, error) {
	var progressList []Progress

	if err := controller.db.Where("status = ?", status).
		Preload("Student").Preload("Title").
		Find(&progressList).Error; err != nil {
		return nil, err
	}
	return progressList, nil
}

func (controller *ProgressController) GetSubmitCountByAssignmentID(assignmentID uint) (uint, error) {
	var count int64

	if err := controller.db.Model(&evalModel.AssignmentSubmission{}).
		Where("assignment_id = ? AND submitted = ?", assignmentID, true).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return uint(count), nil
}
