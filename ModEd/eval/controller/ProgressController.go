//MEP-1006

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
		Preload("Assignment")

	if filters != nil {
		if assignmentId, exists := filters["assignment_id"]; exists {
			query = query.Where("assignment_id = ?", assignmentId)
		}
		if studentCode, exists := filters["student_code"]; exists {
			query = query.Where("student_code = ?", studentCode)
		}
		if submitted, exists := filters["is_submitted"]; exists {
			query = query.Where("submitted = ?", submitted)
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

func (controller *ProgressController) GetProgressByStudentCode(assignmentId uint, studentCode string) ([]evalModel.Progress, error) {
	return controller.List(map[string]interface{}{
		"assignment_id": assignmentId,
		"student_code":  studentCode,
	})
}

func (controller *ProgressController) GetProgressByStatus(assignmentId uint, submitted bool) ([]evalModel.Progress, error) {
	return controller.List(map[string]interface{}{
		"assignment_id": assignmentId,
		"submitted":     submitted,
	})
}

func (controller *ProgressController) GetAssignmentSubmitCount(assignmentId uint) (map[bool]int, error) {
	progressList, err := controller.List(map[string]interface{}{
		"assignment_id": assignmentId,
	})
	if err != nil {
		return nil, err
	}

	statusCount := make(map[bool]int)
	for _, progress := range progressList {
		statusCount[progress.IsSubmitted]++
	}

	return statusCount, nil
}
