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
			ID:         base.ID,
			QuestionDetail: base.QuestionDetail,
			CorrectAnswer:  base.CorrectAnswer,
			Score:          base.Score,
		},
		Choice: f.Choice,
	}
}