package model

type ShortAnswer struct {
	BaseAnswer
	AnswerText string `json:"answer_text"`
}

type ShortAnswerFactory struct {
	AnswerText string
}

func (f ShortAnswerFactory) NewAnswer(questionID, submissionID uint) AnswerProductInterface {
	return &ShortAnswer{
		BaseAnswer: BaseAnswer{
			QuestionID:   questionID,
			SubmissionID: submissionID,
		},
		AnswerText: f.AnswerText,
	}
}
