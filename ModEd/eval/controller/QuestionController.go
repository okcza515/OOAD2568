package controller

import (
	"ModEd/core"
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type IQuestionController interface {
	CreateQuestion(question *model.Question) error
	UpdateQuestion(id uint, updatedQuestion *model.Question) error
	GetQuestionsByExamID(examID uint) ([]model.Question, error)
	DeleteQuestion(id uint) error
}

type QuestionController struct {
	db   *gorm.DB
	core *core.BaseController[*model.Question]
}

func NewQuestionController(db *gorm.DB) *QuestionController {
	return &QuestionController{
		db:   db,
		core: core.NewBaseController[*model.Question](db),
	}
}

func (c *QuestionController) CreateQuestion(question *model.Question) error {
	if err := c.core.Insert(question); err != nil {
		return err
	}
	return nil
}

func (c *QuestionController) UpdateQuestion(id uint, updatedQuestion *model.Question) error {

	if err := c.core.UpdateByCondition(map[string]interface{}{"id": id}, updatedQuestion); err != nil {
		return err
	}
	return nil
}

func (c *QuestionController) GetQuestionsByExamID(examID uint) ([]model.Question, error) {
	var questions []model.Question

	err := c.db.
		Joins("JOIN exam_sections ON exam_sections.id = questions.section_id").
		Where("exam_sections.exam_id = ?", examID).
		Preload("Section").
		Find(&questions).Error

	if err != nil {
		return nil, err
	}

	return questions, nil
}

func (c *QuestionController) DeleteQuestion(id uint) error {
	if err := c.core.DeleteByCondition(map[string]interface{}{"id": id}); err != nil {
		return err
	}
	return nil
}
