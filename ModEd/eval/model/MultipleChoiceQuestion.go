package model

type MultipleChoiceQuestion struct {
	Question
	Choice []string `json:"choices"`
}

type MultipleChoiceQuestionFactory struct {
	Choice []string
}

func (f MultipleChoiceQuestionFactory) CreateQuestion(base Question) QuestionProductInterface {
	return &MultipleChoiceQuestion{
		Question: Question{
			ExamID:         base.ExamID,
			QuestionDetail: base.QuestionDetail,
			QuestionType:   base.QuestionType,
			CorrectAnswer:  base.CorrectAnswer,
			Score:          base.Score,
		},
		Choice: f.Choice,
	}
}