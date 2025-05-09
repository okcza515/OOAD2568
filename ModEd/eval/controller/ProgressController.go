package controller

import (
	"ModEd/core"

	evalModel "ModEd/eval/model"

	"gorm.io/gorm"
)

type ProgressController struct {
	*core.BaseController[evalModel.Progress]
	db *gorm.DB
}

func NewProgressController(db *gorm.DB) *ProgressController {
	return &ProgressController{
		db:             db,
		BaseController: core.NewBaseController[evalModel.Progress](db),
	}
}

func (controller *ProgressController) GetProgress(filter ProgressFilter) ([]evalModel.Progress, error) {
	var ProgressList []evalModel.Progress

	query := controller.db.Model(&evalModel.Progress{}).
		Preload("Student").
		Preload("Assessment")

	if filter.AssessmentId != 0 {
		query = query.Where("AssessmentId = ?", filter.AssessmentId)
	}

	if filter.Type != "" {
		query = query.Where("type = ?", filter.Type)
	}

	if filter.StudentCode != "" {
		query = query.Where("student_code = ?", filter.StudentCode)
	}

	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}

	if err := query.Find(&ProgressList).Error; err != nil {
		return nil, err
	}

	return ProgressList, nil
}

func (controller *ProgressController) GetAllProgress() ([]evalModel.Progress, error) {
	return controller.List(nil, "Student", "Assessment")
}

func (controller *ProgressController) GetProgressByStudentCode(id uint) (evalModel.Progress, error) {
	return controller.RetrieveByID(id, "Student", "Assessment")
}

func (controller *ProgressController) GetProgressByStudentStatus(id uint, Status string) (evalModel.Progress, error) {
	return controller.RetrieveByCondition(map[string]interface{}{
		"id":     id,
		"status": Status,
	}, "Student", "Assessment")
}

func (controller *ProgressController) GetSubmitCount(assessmentId uint) (uint, error) {
	var count int64
	if err := controller.db.Model(&evalModel.AssessmentSubmission{}).
		Where("assessment_id = ? AND submitted = ?", assessmentId, true).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return uint(count), nil
}
