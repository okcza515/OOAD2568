// MEP-1007
package model

type SubmissionFactory interface{}

func NewSubmissionByFactory(questionType QuestionType) SubmissionFactory {
	switch questionType {
	case MultipleChoiceQuestion:
		return &MultipleChoiceAnswerSubmission{}
	case ShortAnswerQuestion:
		return &ShortAnswerSubmission{}
	case TrueFalseQuestion:
		return &TrueFalseAnswerSubmission{}
	default:
		return nil
	}
}
