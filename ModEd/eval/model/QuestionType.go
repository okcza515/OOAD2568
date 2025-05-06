package model

type QuestionType string

const (
	MultipleChoiceQuestionType QuestionType = "MultipleChoiceQuestion"
	ShortAnswerQuestionType    QuestionType = "ShortAnswerQuestion"
	TrueFalseQuestionType      QuestionType = "TrueFalseQuestion"
	SubjectiveQuestionType     QuestionType = "SubjectiveQuestion"
)
