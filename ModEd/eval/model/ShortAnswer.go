package model

import "ModEd/core"

type ShortAnswer struct {
	core.BaseModel
	QuestionID  	uint		`gorm:"type:integer;not null" json:"question_id" csv:"question_id"`
	Question  		Question 	`gorm:"foreignKey:QuestionID;references:ID" json:"question" csv:"question"`
	ExpectedAnswer 	string		`gorm:"type:text;not null" json:"expected_answer" csv:"expected_answer"`
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
