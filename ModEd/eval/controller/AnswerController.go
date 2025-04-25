// MEP-1007
package controller

import (
	"ModEd/eval/model"
	"errors"

	"gorm.io/gorm"
)

type IAnswerController interface {
	SubmitAnswer(answer *model.Answer) error
	GetAnswersByQuestion(questionID uint) ([]model.Answer, error)
	GetAnswersByStudent(studentID uint) ([]model.Answer, error)
	GetAnswerByQuestionAndStudent(questionID uint, studentID uint) (*model.Answer, error)
	UpdateAnswerByID(answerID uint, updatedData map[string]interface{}) error
	DeleteAnswerByID(answerID uint) error
}

type AnswerController struct {
	db *gorm.DB
}

func NewAnswerController(db *gorm.DB) *AnswerController {
	return &AnswerController{db: db}
}

func (c *AnswerController) SubmitAnswer(questionID, studentID uint, answerText string) error {
	var existingAnswer model.Answer

	if err := c.db.Where("question_id = ? AND student_id = ?", questionID, studentID).First(&existingAnswer).Error; err == nil {
		return errors.New("answer already submitted")
	}

	newAnswer := model.Answer{
		QuestionID: questionID,
		StudentID:  studentID,
		Answer:     answerText,
	}

	if err := c.db.Create(&newAnswer).Error; err != nil {
		return err
	}

	return nil
}

func (c *AnswerController) GetAnswersByQuestion(questionID uint) ([]model.Answer, error) {
	var answers []model.Answer
	if err := c.db.Where("question_id = ?", questionID).Preload("Question").Preload("Student").Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}

func (c *AnswerController) GetAnswerByQuestionAndStudent(questionID uint, studentID uint) (*model.Answer, error) {
	var answer model.Answer
	if err := c.db.Where("question_id = ? AND student_id = ?", questionID, studentID).
		Preload("Question").
		Preload("Student").
		First(&answer).Error; err != nil {
		return nil, err
	}
	return &answer, nil
}

func (c *AnswerController) UpdateAnswerByID(answerID uint, updatedData map[string]interface{}) error {
	var existingAnswer model.Answer
	if err := c.db.Where("id = ?", answerID).First(&existingAnswer).Error; err != nil {
		return errors.New("answer not found")
	}

	if err := c.db.Model(&existingAnswer).Updates(updatedData).Error; err != nil {
		return err
	}
	return nil
}

func (c *AnswerController) DeleteAnswerByID(answerID uint) error {
	if err := c.db.Delete(&model.Answer{}, answerID).Error; err != nil {
		return err
	}
	return nil
}
