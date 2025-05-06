package model

import "fmt"

type AnswerProductInterface interface {
	GetQuestionID() uint
	GetSubmissionID() uint
}

type AnswerFactory interface {
	NewAnswer(questionID, submissionID uint) AnswerProductInterface
}
type AnswerFactoryImpl struct {
	Choices   []string
	ShortText string
	Content   string
	Boolean   bool
}

func (af AnswerFactoryImpl) NewAnswer(questionType QuestionType, questionID, submissionID uint, answerData interface{}) (AnswerProductInterface, error) {
	creator, ok := creatorMap[questionType]
	if !ok {
		return nil, fmt.Errorf("unknown question type: %s", questionType)
	}
	return creator(questionID, submissionID, answerData)
}
