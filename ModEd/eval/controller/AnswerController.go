package controller

import (
	"ModEd/eval/model"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type IAnswerController interface {
	SubmitAnswers(factories []model.AnswerFactory, questionIDs []uint, studentID uint, examID uint) ([]model.AnswerProductInterface, error)
}

type AnswerController struct {
	db *gorm.DB
}

func NewAnswerController(db *gorm.DB) *AnswerController {
	return &AnswerController{db: db}
}

func (ac *AnswerController) NewAnswerFactory(questionType string, answerData interface{}) (model.AnswerFactory, error) {
	switch questionType {
	case "multiple_choice":
		choices, ok := answerData.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid data for multiple choice question")
		}
		return model.MCAnswerFactory{Choices: choices}, nil

	case "short_answer":
		answer, ok := answerData.(string)
		if !ok {
			return nil, fmt.Errorf("invalid data for short answer question")
		}
		return model.ShortAnswerFactory{AnswerText: answer}, nil

	case "true_false":
		answer, ok := answerData.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid data for true/false question")
		}
		return model.TrueFalseFactory{Answer: answer}, nil

	case "subjective":
		answer, ok := answerData.(string)
		if !ok {
			return nil, fmt.Errorf("invalid data for subjective question")
		}
		return model.SubjectiveAnswerFactory{Content: answer}, nil

	default:
		return nil, fmt.Errorf("unknown question type: %s", questionType)
	}
}

func (ac *AnswerController) SubmitAnswers(factories []model.AnswerFactory, questionIDs []uint, studentID uint, examID uint) ([]model.AnswerProductInterface, error) {
	var results []model.AnswerProductInterface

	if len(questionIDs) != len(factories) {
		return nil, fmt.Errorf("number of question IDs doesn't match the number of answer factories")
	}

	// var submission model.Answer
	// if err := ac.db.Where("student_id = ? AND exam_id = ?", studentID, examID).First(&submission).Error; err == nil {
	// 	if submission.IsLocked {
	// 		return nil, fmt.Errorf("this exam has already been submitted by this student")
	// 	}
	// }

	submissionTime := time.Now()

	for i, factory := range factories {
		answer := factory.NewAnswer(questionIDs[i], studentID)
		results = append(results, answer)
	}

	isLocked := true

	submission := model.Answer{
		StudentID:   studentID,
		ExamID:      examID,
		SubmittedAt: submissionTime,
		IsLocked:    isLocked,
	}

	if err := ac.db.Create(&submission).Error; err != nil {
		return nil, fmt.Errorf("failed to save submission: %v", err)
	}

	return results, nil
}
