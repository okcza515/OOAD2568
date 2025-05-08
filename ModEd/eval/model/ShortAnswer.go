package model

import "ModEd/core"

type ShortAnswer struct {
	core.BaseModel
	QuestionID     uint
	ExpectedAnswer string
}

// type ShortAnswerQuestionFactory struct{
// 	AnswerText string
// }

// func (f ShortAnswerQuestionFactory) CreateQuestion(base Question) QuestionProductInterface {
// 	return &ShortAnswerQuestion{
// 		Question: Question{
// 			ID:         	base.ID,
// 			QuestionDetail: base.QuestionDetail,
// 			CorrectAnswer:  base.CorrectAnswer,
// 			Score:          base.Score,
// 		},
// 		AnswerText: f.AnswerText,
// 	}
// }
