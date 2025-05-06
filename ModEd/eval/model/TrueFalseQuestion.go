package model

type TrueFalseQuestion struct {
	Question
	Answer bool `json:"answer"`
}

type TrueFalseQuestionFactory struct{
	Answer bool
}

func (f TrueFalseQuestionFactory) CreateQuestion(base Question) QuestionProductInterface {
	return &TrueFalseQuestion{
		Question: Question{
			ExamID:         base.ExamID,
			QuestionDetail: base.QuestionDetail,
			QuestionType:   base.QuestionType,
			CorrectAnswer:  base.CorrectAnswer,
			Score:          base.Score,
		},
		Answer: f.Answer,
	}
}