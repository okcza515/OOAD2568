package controller

import (
	"ModEd/eval/model"

	"gorm.io/gorm"
)

type IQuizController interface {
	ListAllQuizzes() ([]model.Quiz, error)
	RetrieveQuiz(id uint) (*model.Quiz, error)
	InsertQuiz(quiz *model.Quiz) error
	UpdateQuiz(quiz *model.Quiz) error
	DeleteQuiz(id uint) error
}

type QuizController struct {
	db *gorm.DB
}

func NewQuizController(db *gorm.DB) IQuizController {
	return &QuizController{db: db}
}

func (c *QuizController) ListAllQuizzes() ([]model.Quiz, error) {
	var quizzes []model.Quiz
	err := c.db.Preload("Submission").Preload("CourseId").Preload("InstructorId").Find(&quizzes).Error
	return quizzes, err
}

func (c *QuizController) RetrieveQuiz(id uint) (*model.Quiz, error) {
	var quiz model.Quiz
	err := c.db.Preload("Submission").Preload("CourseId").Preload("InstructorId").First(&quiz, id).Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (c *QuizController) InsertQuiz(quiz *model.Quiz) error {
	return c.db.Create(quiz).Error
}

func (c *QuizController) UpdateQuiz(quiz *model.Quiz) error {
	return c.db.Save(quiz).Error
}

func (c *QuizController) DeleteQuiz(id uint) error {
	return c.db.Delete(&model.Quiz{}, id).Error
}
