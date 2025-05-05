package model

type MultipleChoiceAnswer struct {
	BaseAnswer
	Choices []string `json:"choices"`
}

type MCAnswerFactory struct {
	Choices []string
}

func (f MCAnswerFactory) NewAnswer(questionID, submissionID uint) AnswerProductInterface {
	return &MultipleChoiceAnswer{
		BaseAnswer: BaseAnswer{
			QuestionID:   questionID,
			SubmissionID: submissionID,
		},
		Choices: f.Choices,
	}
}
