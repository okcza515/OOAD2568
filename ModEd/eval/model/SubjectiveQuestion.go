package model

type SubjectiveQuestion struct {
	Question
	Content string `json:"content"`
}

type SubjectiveQuestionFactory struct{
	Content string
}

func (f SubjectiveQuestionFactory) CreateQuestion(base Question) QuestionProductInterface {
	return &SubjectiveQuestion{
		Question: Question{
			ID:         	base.ID,
			QuestionDetail: base.QuestionDetail,
			CorrectAnswer:  base.CorrectAnswer,
			Score:          base.Score,
		},
		Content: f.Content,
	}
}