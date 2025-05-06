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
	case "MultipleChoiceQuestion":
		choices, ok := answerData.([]string)
		if !ok {
			return nil, fmt.Errorf("invalid data for multiple choice question")
		}
		return model.MCAnswerFactory{Choices: choices}, nil

	case "ShortAnswerQuestion":
		answer, ok := answerData.(string)
		if !ok {
			return nil, fmt.Errorf("invalid data for short answer question")
		}
		return model.ShortAnswerFactory{AnswerText: answer}, nil

	case "TrueFalseQuestion":
		answer, ok := answerData.(bool)
		if !ok {
			return nil, fmt.Errorf("invalid data for true/false question")
		}
		return model.TrueFalseFactory{Answer: answer}, nil

	case "SubjectiveQuestion":
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

	submissionTime := time.Now()

	if len(questionIDs) != len(factories) {
		return nil, fmt.Errorf("number of question IDs doesn't match the number of answer factories")
	}

	var exam model.Examination
	if err := ac.db.First(&exam, examID).Error; err != nil {
		return nil, fmt.Errorf("exam not found")
	}

	var count int64
	ac.db.Model(&model.Answer{}).Where("student_id = ? AND exam_id = ?", studentID, examID).Count(&count)

	if int(count) >= exam.Attempt {
		return nil, fmt.Errorf("submission limit reached: allowed %d times", exam.Attempt)
	}

	for i, factory := range factories {
		answer := factory.NewAnswer(questionIDs[i], studentID)
		results = append(results, answer)
	}

	submission := model.Answer{
		StudentID:   studentID,
		ExamID:      examID,
		SubmittedAt: submissionTime,
	}

	if err := ac.db.Create(&submission).Error; err != nil {
		return nil, fmt.Errorf("failed to save submission: %v", err)
	}

	return results, nil
}

func (ac *AnswerController) StartExam(studentID uint, examID uint) error {
	startTime := time.Now()

	answer := model.Answer{
		StudentID: studentID,
		ExamID:    examID,
		StartTime: startTime,
	}

	return ac.db.Create(&answer).Error
}

func (ac *AnswerController) GetDuration(submissionID uint) (time.Duration, error) {
	var answer model.Answer

	if err := ac.db.First(&answer, submissionID).Error; err != nil {
		return 0, fmt.Errorf("submission not found: %v", err)
	}

	if answer.StartTime.IsZero() || answer.SubmittedAt.IsZero() {
		return 0, fmt.Errorf("submission has incomplete timestamps")
	}

	duration := answer.SubmittedAt.Sub(answer.StartTime)
	return duration, nil
}
