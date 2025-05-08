package model

import "ModEd/core"

type TrueFalseAnswer struct {
	core.BaseModel
	QuestionID  uint		`gorm:"type:integer;not null" json:"question_id" csv:"question_id"`
	Question  	Question 	`gorm:"foreignKey:QuestionID;references:ID" json:"question" csv:"question"`
	IsExpected  bool		`gorm:"type:boolean;not null" json:"is_expected" csv:"is_expected"`
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
