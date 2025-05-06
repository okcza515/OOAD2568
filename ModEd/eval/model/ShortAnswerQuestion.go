package model

type ShortAnswerQuestion struct {
	Question
	AnswerText string `json:"answer_text"`
}

type ShortAnswerQuestionFactory struct{
	AnswerText string
}

func (f ShortAnswerQuestionFactory) CreateQuestion(base Question) QuestionProductInterface {
	return &ShortAnswerQuestion{
		Question: Question{
			ID:         	base.ID,
			QuestionDetail: base.QuestionDetail,
			CorrectAnswer:  base.CorrectAnswer,
			Score:          base.Score,
		},
		AnswerText: f.AnswerText,
	}
}