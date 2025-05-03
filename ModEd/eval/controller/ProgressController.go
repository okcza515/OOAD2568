package controller

import (
	commonModel "ModEd/common/model"

	evalModel "ModEd/eval/model"

	"time"

	"gorm.io/gorm"
)

type ProgressSearchStrategy interface {
	Search(db *gorm.DB) ([]Progress, error)
}

type BaseProgressSearchStrategy struct {
	AssignmentId uint
	QuizId       uint
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
	Status string
}

type Progress struct {
	gorm.Model
	StudentCode      commonModel.Student  `gorm:"foreignKey:StudentCode;references:StudentCode"`
	Title            evalModel.Assignment `gorm:"foreignKey:AssignmentId;references:AssignmentId"`
	AssignmentId     evalModel.Assignment `gorm:"foreignKey:AssignmentId;references:AssignmentId"`
	QuizId           evalModel.Quiz       `gorm:"foreignKey:QuizId;references:QuizId"`
	AssignmentStatus evalModel.Assignment `gorm:"foreignKey:AssignmentId;references:AssignmentId"`
	QuizStatus       evalModel.Quiz       `gorm:"foreignKey:QuizId;references:QuizId"`
	LastUpdate       time.Time            `gorm:"autoUpdateTime"`
	TotalSubmit      uint
}

type ProgressController struct {
	db *gorm.DB
}

func NewProgressController(db *gorm.DB) *ProgressController {
	return &ProgressController{db: db}
}

func (strategy *GetAllStudentProgressStrategy) Search(db *gorm.DB) ([]Progress, error) {
	var progressList []Progress
	query := db.Model(&Progress{})

	if strategy.AssignmentId != 0 {
		query = query.Where("assignment_id = ?", strategy.AssignmentId)
	}
	if strategy.QuizId != 0 {
		query = query.Where("quiz_id = ?", strategy.QuizId)
	}

	if err := query.Preload("Student").
		Preload("Title").
		Preload("AssignmentStatus").
		Preload("QuizStatus").
		Find(&progressList).Error; err != nil {
		return nil, err
	}
	return progressList, nil
}

func (strategy *GetProgressByStudentCodeStrategy) Search(db *gorm.DB) ([]Progress, error) {
	var progressList []Progress
	query := db.Model(&Progress{})

	if strategy.AssignmentId != 0 {
		query = query.Where("assignment_id = ?", strategy.AssignmentId)
	}
	if strategy.QuizId != 0 {
		query = query.Where("quiz_id = ?", strategy.QuizId)
	}

	if err := query.Where("student_code = ?", strategy.StudentCode).
		Preload("Student").
		Preload("Title").
		Preload("AssignmentStatus").
		Preload("QuizStatus").
		Find(&progressList).Error; err != nil {
		return nil, err
	}
	return progressList, nil
}

func (strategy *GetProgressByStatusStrategy) Search(db *gorm.DB) ([]Progress, error) {
	var progressList []Progress
	query := db.Model(&Progress{})

	if strategy.AssignmentId != 0 {
		query = query.Where("assignment_status = ?", strategy.Status)
	}
	if strategy.QuizId != 0 {
		query = query.Where("quiz_status = ?", strategy.Status)
	}

	if err := query.Preload("Student").
		Preload("Title").
		Preload("AssignmentStatus").
		Preload("QuizStatus").
		Find(&progressList).Error; err != nil {
		return nil, err
	}
	return progressList, nil
}

func (controller *ProgressController) GetAllProgress(AssignmentId uint, QuizId uint) ([]Progress, error) {
	strategy := &GetAllStudentProgressStrategy{
		BaseProgressSearchStrategy: BaseProgressSearchStrategy{
			AssignmentId: AssignmentId,
			QuizId:       QuizId,
		},
	}
	return strategy.Search(controller.db)
}

func (controller *ProgressController) GetProgressByStudentCode(AssignmentId uint, QuizId uint, StudentCode string) ([]Progress, error) {
	strategy := &GetProgressByStudentCodeStrategy{
		BaseProgressSearchStrategy: BaseProgressSearchStrategy{
			AssignmentId: AssignmentId,
			QuizId:       QuizId,
		},
		StudentCode: StudentCode,
	}
	return strategy.Search(controller.db)
}

func (controller *ProgressController) GetProgressByStatus(AssignmentId uint, QuizId uint, status string) ([]Progress, error) {
	strategy := &GetProgressByStatusStrategy{
		BaseProgressSearchStrategy: BaseProgressSearchStrategy{
			AssignmentId: AssignmentId,
			QuizId:       QuizId,
		},
		Status: status,
	}
	return strategy.Search(controller.db)
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

func (controller *ProgressController) GetSubmitCountByQuizID(quizId uint) (uint, error) {
	var count int64
	if err := controller.db.Model(&evalModel.QuizSubmission{}).
		Where("quiz_id = ? AND submitted = ?", quizId, true).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return uint(count), nil
}
