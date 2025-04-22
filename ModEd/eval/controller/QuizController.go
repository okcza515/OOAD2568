package controller

import (
	"errors"
	"time"

	quizModel "ModEd/eval/model"

	"gorm.io/gorm"
)

type QuizController interface {
	GetAll() ([]quizModel.Quiz, error)
	GetByID(id uint) (*quizModel.Quiz, error)
	Create(input QuizInput) (*quizModel.Quiz, error)
	Update(id uint, input QuizInput) (*quizModel.Quiz, error)
	Delete(id uint) error
}

type quizController struct {
	db *gorm.DB
}

func NewQuizController(db *gorm.DB) QuizController {
	return &quizController{db: db}
}

type QuizInput struct {
	Title       string
	Description string
	Released    bool
	QuizStart   time.Time
	QuizEnd     time.Time
	Status      string
}

func (c *quizController) GetAll() ([]quizModel.Quiz, error) {
	var quizzes []quizModel.Quiz
	if err := c.db.Preload("Submission").Find(&quizzes).Error; err != nil {
		return nil, err
	}
	return quizzes, nil
}

func (c *quizController) GetByID(id uint) (*quizModel.Quiz, error) {
	var quiz quizModel.Quiz
	if err := c.db.Preload("Submission").First(&quiz, id).Error; err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (c *quizController) Create(input QuizInput) (*quizModel.Quiz, error) {
	quiz := quizModel.Quiz{
		Title:       input.Title,
		Description: input.Description,
		Released:    input.Released,
		QuizStart:   input.QuizStart,
		QuizEnd:     input.QuizEnd,
		Status:      input.Status,
	}
	if err := c.db.Create(&quiz).Error; err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (c *quizController) Update(id uint, input QuizInput) (*quizModel.Quiz, error) {
	var quiz quizModel.Quiz
	if err := c.db.First(&quiz, id).Error; err != nil {
		return nil, err
	}

	quiz.Title = input.Title
	quiz.Description = input.Description
	quiz.Released = input.Released
	quiz.QuizStart = input.QuizStart
	quiz.QuizEnd = input.QuizEnd
	quiz.Status = input.Status

	if err := c.db.Save(&quiz).Error; err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (c *quizController) Delete(id uint) error {
	var quiz quizModel.Quiz
	if err := c.db.First(&quiz, id).Error; err != nil {
		return errors.New("quiz not found")
	}
	return c.db.Delete(&quiz).Error
}
