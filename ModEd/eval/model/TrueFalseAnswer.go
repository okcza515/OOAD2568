package model

import "ModEd/core"

type TrueFalseAnswer struct {
	core.BaseModel
	QuestionID uint
	IsExpected bool
}

// type TrueFalseQuestionFactory struct{
// 	Answer bool
// }

// func (f TrueFalseQuestionFactory) CreateQuestion(base Question) QuestionProductInterface {
// 	return &TrueFalseQuestion{
// 		Question: Question{
// 			ID:         	base.ID,
// 			QuestionDetail: base.QuestionDetail,
// 			CorrectAnswer:  base.CorrectAnswer,
// 			Score:          base.Score,
// 		},
// 		Answer: f.Answer,
// 	}
// }
