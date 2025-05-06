package controller

import (
	commonModel "ModEd/common/model"

	evalModel "ModEd/eval/model"

	"fmt"

	"time"

	"gorm.io/gorm"
)

type ProgressSearchStrategy interface {
	Search(db *gorm.DB) ([]Progress, error)
}

type BaseProgressSearchStrategy struct {
	AssessmentId uint
	Type         evalModel.AssessmentType
}

type GetAllStudentProgressStrategy struct {
	BaseProgressSearchStrategy
}

type GetProgressByStudentCodeStrategy struct {
	BaseProgressSearchStrategy
	StudentCode string
}

type GetProgressByStatusStrategy struct {
	BaseProgressSearchStrategy
	Status evalModel.AssessmentStatus
}

type Progress struct {
	gorm.Model
	StudentCode  commonModel.Student        `gorm:"foreignKey:StudentCode;references:StudentCode"`
	AssessmentId evalModel.Assessment       `gorm:"foreignKey:AssessmentId;references:AssessmentId"`
	Type         evalModel.AssessmentType   `gorm:"foreignKey:AssessmentId;references:AssessmentId"`
	Status       evalModel.AssessmentStatus `gorm:"foreignKey:AssessmentId;references:AssessmentId"`
	LastUpdate   time.Time                  `gorm:"autoUpdateTime"`
	TotalSubmit  uint
}

type ProgressController struct {
	db *gorm.DB
}

func NewProgressController(db *gorm.DB) *ProgressController {
	return &ProgressController{db: db}
}

func (controller *ProgressController) GetAssessmentsByType(AssessmentType evalModel.AssessmentType) ([]evalModel.Assessment, error) {
	var Assessments []evalModel.Assessment
	if err := controller.db.Where("type = ?", AssessmentType).Find(&Assessments).Error; err != nil {
		return nil, fmt.Errorf("error getting %s list: %v", AssessmentType, err)
	}
	return Assessments, nil
}

func (strategy *GetAllStudentProgressStrategy) Search(db *gorm.DB) ([]Progress, error) {
	var progressList []Progress
	query := db.Model(&Progress{})

	if strategy.AssessmentId != 0 {
		query = query.Where("assessment_id = ?", strategy.AssessmentId)
	}

	if strategy.Type != "" {
		query = query.Where("type = ?", strategy.Type)
	}

	if err := query.Preload("Student").
		Preload("Assessment").
		Find(&progressList).Error; err != nil {
		return nil, err
	}
	return progressList, nil
}

func (strategy *GetProgressByStudentCodeStrategy) Search(db *gorm.DB) ([]Progress, error) {
	var progressList []Progress
	query := db.Model(&Progress{})

	if strategy.AssessmentId != 0 {
		query = query.Where("assessment_id = ?", strategy.AssessmentId)
	}

	if strategy.Type != "" {
		query = query.Where("type = ?", strategy.Type)
	}

	if err := query.Where("student_code = ?", strategy.StudentCode).
		Preload("Student").
		Preload("Assessment").
		Find(&progressList).Error; err != nil {
		return nil, err
	}
	return progressList, nil
}

func (strategy *GetProgressByStatusStrategy) Search(db *gorm.DB) ([]Progress, error) {
	var progressList []Progress
	query := db.Model(&Progress{})

	if strategy.AssessmentId != 0 {
		query = query.Where("assessment_id = ?", strategy.AssessmentId)
	}

	if strategy.Type != "" {
		query = query.Where("type = ?", strategy.Type)
	}

	if err := query.Where("status = ?", strategy.Status).
		Preload("Student").
		Preload("Assessment").
		Find(&progressList).Error; err != nil {
		return nil, err
	}
	return progressList, nil
}

func (controller *ProgressController) GetAllProgressByType(AssessmentType evalModel.AssessmentType, AssessmentId uint) ([]Progress, error) {
	if AssessmentType != evalModel.QuizType && AssessmentType != evalModel.AssignmentType {
		return nil, fmt.Errorf("invalid assessment type: %s", AssessmentType)
	}

	strategy := &GetAllStudentProgressStrategy{
		BaseProgressSearchStrategy: BaseProgressSearchStrategy{
			AssessmentId: AssessmentId,
			Type:         AssessmentType,
		},
	}
	return strategy.Search(controller.db)
}

func (controller *ProgressController) GetProgressByStudentCode(AssessmentType evalModel.AssessmentType, AssessmentId uint, StudentCode string) ([]Progress, error) {
	if AssessmentType != evalModel.QuizType && AssessmentType != evalModel.AssignmentType {
		return nil, fmt.Errorf("invalid assessment type: %s", AssessmentType)
	}

	strategy := &GetProgressByStudentCodeStrategy{
		BaseProgressSearchStrategy: BaseProgressSearchStrategy{
			AssessmentId: AssessmentId,
			Type:         AssessmentType,
		},
		StudentCode: StudentCode,
	}
	return strategy.Search(controller.db)
}

func (controller *ProgressController) GetProgressByStatus(AssessmentType evalModel.AssessmentType, AssessmentId uint, Status evalModel.AssessmentStatus) ([]Progress, error) {
	if AssessmentType != evalModel.QuizType && AssessmentType != evalModel.AssignmentType {
		return nil, fmt.Errorf("invalid assessment type: %s", AssessmentType)
	}

	strategy := &GetProgressByStatusStrategy{
		BaseProgressSearchStrategy: BaseProgressSearchStrategy{
			AssessmentId: AssessmentId,
			Type:         AssessmentType,
		},
		Status: Status,
	}
	return strategy.Search(controller.db)
}

func (controller *ProgressController) GetSubmitCount(AssessmentId uint) (uint, error) {
	var count int64
	if err := controller.db.Model(&evalModel.AssessmentSubmission{}).
		Where("assessment_id = ? AND submitted = ?", AssessmentId, true).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return uint(count), nil
}
