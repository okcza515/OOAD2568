// MEP-1007
package model

type QuestionType string

const (
	MultipleChoiceQuestion QuestionType = "MultipleChoiceQuestion"
	ShortAnswerQuestion    QuestionType = "ShortAnswerQuestion"
	TrueFalseQuestion      QuestionType = "TrueFalseQuestion"
)
