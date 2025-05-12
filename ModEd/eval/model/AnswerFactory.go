// MEP-1007
package model

type AnswerFactory interface{}

func NewAnswerByFactory(questionType QuestionType) AnswerFactory {
	switch questionType {
	case MultipleChoiceQuestion:
		return &MultipleChoiceAnswer{}
	case ShortAnswerQuestion:
		return &ShortAnswer{}
	case TrueFalseQuestion:
		return &TrueFalseAnswer{}
	default:
		return nil
	}
}
