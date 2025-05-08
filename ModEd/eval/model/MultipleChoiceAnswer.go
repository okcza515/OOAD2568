package model

import "ModEd/core"

type MultipleChoiceAnswer struct {
	core.BaseModel
	QuestionID  uint
	AnswerLabel string
	IsExpected  bool
}

// type MultipleChoiceQuestionFactory struct {
// 	Choice []string
// }

// func (f MultipleChoiceQuestionFactory) CreateQuestion(base Question) QuestionProductInterface {
// 	return &MultipleChoiceQuestion{
// 		Question: Question{
// 			ID:         base.ID,
// 			QuestionDetail: base.QuestionDetail,
// 			CorrectAnswer:  base.CorrectAnswer,
// 			Score:          base.Score,
// 		},
// 		Choice: f.Choice,
// 	}
// }
