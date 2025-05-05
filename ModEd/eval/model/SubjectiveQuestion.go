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
			ExamID:         base.ExamID,
			QuestionDetail: base.QuestionDetail,
			QuestionType:   base.QuestionType,
			CorrectAnswer:  base.CorrectAnswer,
			Score:          base.Score,
		},
		Content: f.Content,
	}
}