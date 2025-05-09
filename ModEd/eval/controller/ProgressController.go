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

func (controller *ProgressController) List(filters map[string]interface{}) ([]evalModel.Progress, error) {
	var progressList []evalModel.Progress

	query := controller.db.Model(&evalModel.Progress{}).
		Preload("Student").
		Preload("Assessment")

	if filters != nil {
		if assessmentId, exists := filters["assessment_id"]; exists {
			query = query.Where("assessment_id = ?", assessmentId)
		}
		if studentCode, exists := filters["student_code"]; exists {
			query = query.Where("student_code = ?", studentCode)
		}
		if status, exists := filters["status"]; exists {
			query = query.Where("status = ?", status)
		}
	}

	if err := query.Find(&progressList).Error; err != nil {
		return nil, err
	}

	return progressList, nil
}

func (controller *ProgressController) GetAllProgress() ([]evalModel.Progress, error) {
	return controller.List(nil)
}

func (controller *ProgressController) GetProgressByStudentCode(assessmentId uint, studentCode string) ([]evalModel.Progress, error) {
	return controller.List(map[string]interface{}{
		"assessment_id": assessmentId,
		"student_code":  studentCode,
	})
}

func (controller *ProgressController) GetProgressByStatus(assessmentId uint, status string) ([]evalModel.Progress, error) {
	return controller.List(map[string]interface{}{
		"assessment_id": assessmentId,
		"status":        status,
	})
}

func (controller *ProgressController) GetAssessmentSubmitCount(assessmentId uint) (map[evalModel.AssessmentStatus]int, error) {
	progressList, err := controller.List(map[string]interface{}{
		"assessment_id": assessmentId,
	})
	if err != nil {
		return nil, err
	}

	statusCount := make(map[evalModel.AssessmentStatus]int)
	for _, progress := range progressList {
		statusCount[progress.Status]++
	}

	return statusCount, nil
}
