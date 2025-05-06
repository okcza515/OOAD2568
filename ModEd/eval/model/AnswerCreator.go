package model

import "fmt"

type AnswerCreatorFunc func(questionID, submissionID uint, data interface{}) (AnswerProductInterface, error)

var creatorMap = map[QuestionType]AnswerCreatorFunc{
	MultipleChoiceQuestionType: createMultipleChoiceAnswer,
	ShortAnswerQuestionType:    createShortAnswer,
	TrueFalseQuestionType:      createTrueFalseAnswer,
	SubjectiveQuestionType:     createSubjectiveAnswer,
}

func createMultipleChoiceAnswer(questionID, submissionID uint, data interface{}) (AnswerProductInterface, error) {
	choices, ok := data.([]string)
	if !ok {
		return nil, fmt.Errorf("invalid data for multiple choice question")
	}
	return &MultipleChoiceAnswer{
		BaseAnswer: BaseAnswer{QuestionID: questionID, SubmissionID: submissionID},
		Choices:    choices,
	}, nil
}

func createShortAnswer(questionID, submissionID uint, data interface{}) (AnswerProductInterface, error) {
	text, ok := data.(string)
	if !ok {
		return nil, fmt.Errorf("invalid data for short answer")
	}
	return &ShortAnswer{
		BaseAnswer: BaseAnswer{QuestionID: questionID, SubmissionID: submissionID},
		ShortText:  text,
	}, nil
}

func createTrueFalseAnswer(questionID, submissionID uint, data interface{}) (AnswerProductInterface, error) {
	val, ok := data.(bool)
	if !ok {
		return nil, fmt.Errorf("invalid data for true/false question")
	}
	return &TrueFalseAnswer{
		BaseAnswer: BaseAnswer{QuestionID: questionID, SubmissionID: submissionID},
		Boolean:    val,
	}, nil
}

func createSubjectiveAnswer(questionID, submissionID uint, data interface{}) (AnswerProductInterface, error) {
	text, ok := data.(string)
	if !ok {
		return nil, fmt.Errorf("invalid data for subjective question")
	}
	return &SubjectiveAnswer{
		BaseAnswer: BaseAnswer{QuestionID: questionID, SubmissionID: submissionID},
		AnswerText: text,
	}, nil
}
