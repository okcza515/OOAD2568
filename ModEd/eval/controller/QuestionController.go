package controller

import (
	model "ModEd/eval/model"
	"gorm.io/gorm"
)

type IQuestionController interface {
	CreateQuestion(question *model.Question) error
	UpdateQuestion(id uint, updatedQuestion *model.Question) error
	CreateQuestion(question *model.Question) error
	UpdateCorrectAnswer(id uint, correctAnswer string) error
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

func (c *QuestionController) CreateQuestion(question *model.Question) error {
	if err := c.db.Create(question).Error; err != nil {
		return err
	}
	return nil
}

func (c *QuestionController) UpdateCorrectAnswer(id uint, correctAnswer string) error {
	var question model.Question
	if err := c.db.First(&question, id).Error; err != nil {
		return err
	}

	if err := c.db.Model(&question).Update("Correct_answer", correctAnswer).Error; err != nil {
		return err
	}

	return nil
}