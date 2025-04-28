package controller

import (
	model "ModEd/eval/model"
	"gorm.io/gorm"
)

type IQuestionController interface {
	CreateQuestion(question *model.Question) error
	UpdateQuestion(id uint, updatedQuestion *model.Question) error
	GetQuestionsByExamID(examID uint) ([]model.Question, error)
}

type QuestionController struct {
	db *gorm.DB
}

func NewQuestionController(db *gorm.DB) *QuestionController {
	return &QuestionController{db: db}
}

func (c *QuestionController) CreateQuestion(question *model.Question) error {
	if err := c.db.Create(question).Error; err != nil {
		return err
	}
	return nil
}

func (c *QuestionController) UpdateQuestion(id uint, updatedQuestion *model.Question) error {
	var question model.Question
	if err := c.db.First(&question, id).Error; err != nil {
		return err
	}

	if err := c.db.Model(&question).Select("*").Updates(updatedQuestion).Error; err != nil {
		return err
	}
	return nil
}

func (c *QuestionController) GetQuestionsByExamID(examID uint) ([]model.Question, error) {
    var questions []model.Question
    if err := c.db.Where("exam_id = ?", examID).Find(&questions).Error; err != nil {
        return nil, err
    }
    return questions, nil
}