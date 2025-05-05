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
			ExamID:         base.ExamID,
			QuestionDetail: base.QuestionDetail,
			QuestionType:   base.QuestionType,
			CorrectAnswer:  base.CorrectAnswer,
			Score:          base.Score,
		},
		AnswerText: f.AnswerText,
	}
}