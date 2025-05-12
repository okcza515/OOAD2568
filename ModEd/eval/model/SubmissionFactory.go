package model

type SubmissionFactory interface {
	SetQuestionID(id uint)
	GetQuestionID() uint
	SetSubmmissionID(id uint)
	GetSubmissionID() uint
}

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
