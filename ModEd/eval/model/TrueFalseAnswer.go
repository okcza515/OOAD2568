package model

type TrueFalse struct {
	BaseAnswer
	Answer bool `json:"answer"`
}

type TrueFalseFactory struct {
	Answer bool
}

func (f TrueFalseFactory) NewAnswer(questionID, submissionID uint) AnswerProductInterface {
	return &TrueFalse{
		BaseAnswer: BaseAnswer{
			QuestionID:   questionID,
			SubmissionID: submissionID,
		},
		Answer: f.Answer,
	}
}
