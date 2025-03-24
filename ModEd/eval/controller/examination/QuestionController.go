package controller

import (
	model "ModEd/eval/model"
	"gorm.io/gorm"
)

type IQuestionController interface {
	CreateQuestion(question *model.Question) error
	UpdateQuestion(id uint, updatedQuestion *model.Question) error
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

	// Update only the fields that have non-zero values
	if err := c.db.Model(&question).Select("*").Updates(updatedQuestion).Error; err != nil {
		return err
	}
	return nil
}
