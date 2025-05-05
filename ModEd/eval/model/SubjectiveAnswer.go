package model

type SubjectiveAnswer struct {
	BaseAnswer
	Content string `json:"content"`
}

type SubjectiveAnswerFactory struct {
	Content string
}

func (f SubjectiveAnswerFactory) NewAnswer(questionID, submissionID uint) AnswerProductInterface {
	return &SubjectiveAnswer{
		BaseAnswer: BaseAnswer{
			QuestionID:   questionID,
			SubmissionID: submissionID,
		},
		Content: f.Content,
	}
}
