package controller

import (
	commonModel "ModEd/common/model"

	"ModEd/core"

	evalModel "ModEd/eval/model"

	"fmt"

	"time"

	"gorm.io/gorm"
)

type Progress struct {
	gorm.Model
	StudentCode  string
	Student      commonModel.Student `gorm:"foreignKey:StudentCode;references:StudentCode"`
	AssessmentId uint
	Assessment   evalModel.Assessment       `gorm:"foreignKey:AssessmentId;references:AssessmentId"`
	Type         evalModel.AssessmentType   `gorm:"column:type;not null"`
	Status       evalModel.AssessmentStatus `gorm:"column:status;not null"`
	LastUpdate   time.Time                  `gorm:"column:last_update;autoUpdateTime"`
	TotalSubmit  uint                       `gorm:"column:total_submit;default:0"`
}

func (p Progress) GetID() uint {
	return p.Model.ID
}

func (p Progress) ToString() string {
	return fmt.Sprintf("Progress{ID: %d, StudentCode: %s, AssessmentId: %d, Type: %s, Status: %s, LastUpdate: %v, TotalSubmit: %d}",
		p.ID, p.StudentCode, p.AssessmentId, p.Type, p.Status, p.LastUpdate, p.TotalSubmit)
}

func (p Progress) Validate() error {
	if p.StudentCode == "" {
		return fmt.Errorf("Student code is required")
	}
	if p.AssessmentId == 0 {
		return fmt.Errorf("Assessment ID is required")
	}
	if p.Type == "" {
		return fmt.Errorf("Assessment type is required")
	}
	if p.Status == "" {
		return fmt.Errorf("Status is required")
	}
	return nil
}

func (p Progress) ToCSVRow() string {
	return fmt.Sprintf("%d, %s, %d, %s, %s, %v, %d",
		p.ID, p.StudentCode, p.AssessmentId, p.Type, p.Status, p.LastUpdate, p.TotalSubmit)
}

func (p Progress) FromCSV(raw string) error {
	// TODO: Implement CSV parsing
	return nil
}

func (p Progress) ToJSON() string {
	return fmt.Sprintf(`{"id":%d,"student_code":"%s","assessment_id":%d,"type":"%s","status":"%s","last_update":"%v","total_submit":%d}`,
		p.ID, p.StudentCode, p.AssessmentId, p.Type, p.Status, p.LastUpdate, p.TotalSubmit)
}

func (p Progress) FromJSON(raw string) error {
	// TODO: Implement JSON parsing
	return nil
}

type ProgressFilter struct {
	AssessmentId uint
	Type         evalModel.AssessmentType
	StudentCode  string
	Status       evalModel.AssessmentStatus
}

type ProgressController struct {
	*core.BaseController[Progress]
	db *gorm.DB
}

func NewProgressController(db *gorm.DB) *ProgressController {
	return &ProgressController{
		db:             db,
		BaseController: core.NewBaseController[Progress](db),
	}
}

func (controller *ProgressController) GetProgress(filter ProgressFilter) ([]Progress, error) {
	var ProgressList []Progress

	query := controller.db.Model(&Progress{}).
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

func (controller *ProgressController) GetProgressByID(id uint) (Progress, error) {
	return controller.RetrieveByID(id, "Student", "Assessment")
}

func (controller *ProgressController) ListAllProgress() ([]Progress, error) {
	return controller.List(nil, "Student", "Assessment")
}

func (controller *ProgressController) ListProgressWithPagination(page, pageSize int) ([]Progress, error) {
	return controller.ListPagination(nil, page, pageSize, "Student", "Assessment")
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
